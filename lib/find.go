package lib

import (
	"fmt"
	"strings"
)

// FindExactIndex locates the exact location of an entry
func (l *LocationDB) FindIndexExact(location string) int {
	for index, candidate := range l.Locations {
		if candidate.Path == location {
			return index
		}
	}
	return -1
}

// FindIndex fuzzily locates a location
func (l *LocationDB) FindIndexFuzzy(location string) int {
	for index, candidate := range l.Locations {
		if strings.Contains(candidate.Path, location) {
			return index
		}
	}
	return -1
}

// FindLocation finds a location inside the LocationDB
func (l *LocationDB) FindLocation(location string) (string, error) {
	index := l.FindIndexFuzzy(location)
	if index < 0 {
		return "", fmt.Errorf("Could not find entry for %s", location)
	}
	return l.Locations[index].Path, nil
}
