package lib

import (
	"fmt"
	"strings"
)

func (l *LocationDB) findIndex(location string) int {
	for index, candidate := range l.Locations {
		if strings.Contains(candidate.Path, location) {
			return index
		}
	}
	return -1
}

// FindLocation finds a location inside the LocationDB
func (l *LocationDB) FindLocation(location string) (string, error) {
	index := l.findIndex(location)
	if index < 0 {
		return "", fmt.Errorf("Could not find entry for %s", location)
	}
	return l.Locations[index].Path, nil
}
