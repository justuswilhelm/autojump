package lib

func (l *LocationDB) appendLocation(location string) {
	l.Locations = append(l.Locations, Location{
		Frequency: 1,
		Path:      location,
	})
}

// RecordLocation creates or updates a visited location
func (l *LocationDB) RecordLocation(location string) error {
	index := l.findIndex(location)
	if index < 0 {
		l.appendLocation(location)
		return nil
	}
	l.Locations[index].Frequency++
	return nil
}
