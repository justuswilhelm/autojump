package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	// DBLocation stores location of autojump DB
	DBLocation = ".config/autojump/db"
)

var (
	location       string
	record         bool
	db             []string
	userDBLocation string
)

func init() {
	flag.StringVar(&location, "location", "", "Where to jump to")
	flag.BoolVar(&record, "record", false, "Whether to record a jump")
	flag.Parse()
}

func ensureDbPath() error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf(
			"Error when retrieving current user: %+v", err)
	}
	dir := usr.HomeDir
	userDBLocation = filepath.Join(dir, DBLocation)

	path := filepath.Dir(userDBLocation)
	if err := os.MkdirAll(path, 040755); err != nil {
		return fmt.Errorf(
			"Error when creating config folder %s: %+v", path, err)
	}
	file, err := os.OpenFile(userDBLocation, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf(
			"Error when creating db file %s: %+v", userDBLocation, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Oh no %+v", err)
		}
	}()
	return nil
}

func readDb() ([]string, error) {
	result, err := ioutil.ReadFile(userDBLocation)
	if err != nil {
		return nil, fmt.Errorf("Error when reading database: %+v", err)
	}
	split := strings.Split(string(result), "\n")
	return split, nil
}

func findLocation(destPath string) (string, error) {
	for _, path := range db {
		if strings.Contains(path, destPath) {
			return path, nil
		}
	}
	return "", fmt.Errorf("Could not find path")
}

func recordLocation(location string) error {
	db = append(db, location)

	file, err := os.OpenFile(userDBLocation, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Oh no %+v", err)
		}
	}()

	buf := fmt.Sprintf("%s\n", location)

	_, err = file.WriteString(buf)
	return err
}

func main() {
	if err := ensureDbPath(); err != nil {
		log.Fatalf("Error when ensuring Db path: %+v", err)
	}
	result, err := readDb()
	if err != nil {
		log.Fatalf("Error when reading db: %+v", err)
	}
	db = result

	if record {
		if err := recordLocation(location); err != nil {
			log.Fatalf("Error when recording new location: %+v", err)
		}
	} else {
		if location == "" {
			log.Fatalf("Must specify location")
		}
		location, err := findLocation(location)
		if err != nil {
			log.Fatalf("Error when searching location: %+v", err)
		}
		fmt.Print(location)
	}
}
