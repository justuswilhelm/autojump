package lib

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func (l *LocationDB) dump(f *os.File) error {
	w := csv.NewWriter(f)

	for index, location := range l.Locations {
		row := []string{fmt.Sprint(location.Frequency), location.Path}
		err := w.Write(row)
		if err != nil {
			return fmt.Errorf(
				"Error when writing database row %d: %+v", index, err)
		}
	}
	w.Flush()
	return nil
}

func ensureFolder() error {
	path, err := GetDBFolder()
	if err != nil {
		return fmt.Errorf("Error when retrieving DB path: %v", err)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return fmt.Errorf("Error when creating folder: %v", err)
		}
	}
	return nil
}

// DumpDB writes the database back to its filename
func (l *LocationDB) DumpDB() error {
	if err := ensureFolder(); err != nil {
		log.Fatalf("Error when ensuring db folder: %+v", err)
	}
	f, err := os.OpenFile(l.FilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Error when opening location db %s", l.FilePath)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Error when closing %s: %+v", l.FilePath, err)
		}
	}()

	if err := l.dump(f); err != nil {
		return fmt.Errorf("Error when dumping db to %s: %+v", l.FilePath, err)
	}

	return nil
}

// EchoDB writes the database to stdout
func (l *LocationDB) EchoDB() error {
	if err := l.dump(os.Stdout); err != nil {
		return fmt.Errorf("Error when dumping db to stdout: %+v", err)
	}
	return nil
}
