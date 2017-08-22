package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/justuswilhelm/autojump/lib"
)

var (
	location string
	record   bool
	echo     bool
)

func run() error {
	path, err := lib.GetDBPath()
	if err != nil {
		return fmt.Errorf("Error when retrieving DB path: %+v", err)
	}
	db, err := lib.ReadDB(path)
	if err != nil {
		return fmt.Errorf("Error when reading db: %+v", err)
	}

	if echo {
		if err := db.EchoDB(); err != nil {
			return fmt.Errorf("Error when dumping db: %+v", err)
		}
		return nil
	}

	if location == "" {
		return fmt.Errorf("Must specify location")
	}

	if record {
		if err := db.RecordLocation(location); err != nil {
			return fmt.Errorf("Error when recording new location: %+v", err)
		}
		if err := db.DumpDB(); err != nil {
			return fmt.Errorf("Error dumping db: %+v", err)
		}
		return nil
	}
	location, err := db.FindLocation(location)
	if err != nil {
		return fmt.Errorf("Error when searching location: %+v", err)
	}
	fmt.Print(location)
	return nil
}

func init() {
	flag.BoolVar(&echo, "echo", false, "Echo DB to stdout")
	flag.BoolVar(&record, "record", false, "Whether to record a jump")
	flag.StringVar(&location, "location", "", "Where to jump to")
	flag.Parse()
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Runtime error: %+v", err)
	}
}
