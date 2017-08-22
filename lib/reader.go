package lib

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

// GetDBPath retrieves DB Path for user.
func GetDBPath() (string, error) {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome != "" {
		return filepath.Join(configHome, DBLocation), nil
	}
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf(
			"Error when retrieving current user: %+v", err)
	}
	dir := usr.HomeDir
	return filepath.Join(dir, ".config", DBLocation), nil
}

// ReadDB reads the database from a given path
func ReadDB(path string) (*LocationDB, error) {
	result := &LocationDB{
		FilePath: path,
	}
	file, err := os.Open(result.FilePath)

	if err, ok := err.(*os.PathError); ok && err.Err == syscall.ENOENT {
		result.Locations = make([]Location, 0)
		return result, nil
	} else if err != nil {
		return nil, fmt.Errorf("Error when opening database: %+v", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error when closing %s: %+v", result.FilePath, err)
		}
	}()

	result.Locations, err = parseLocations(file)
	if err != nil {
		return nil, fmt.Errorf("Error when parsing database: %+v", err)
	}
	return result, nil
}

func parseLocations(file *os.File) ([]Location, error) {
	reader := csv.NewReader(file)
	raw, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error when reading records: %+v", err)
	}

	result := make([]Location, len(raw))
	for i, r := range raw {
		frequency, err := strconv.ParseUint(r[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf(
				"Error decoding frequency in row %d: %+v",
				i,
				err,
			)
		}
		path := r[1]
		result[i] = Location{
			Frequency: frequency,
			Path:      path,
		}
	}
	return result, nil
}
