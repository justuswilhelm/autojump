package lib

const (
	// DBLocation stores location of autojump DB
	DBFolder = "autojump"
	DBName   = "db"
)

// Location stores a path and the corresponding frequency
type Location struct {
	Path      string
	Frequency uint64
}

// LocationDB stores a number of locations that have been visited and
// the frequency count
type LocationDB struct {
	FilePath  string
	Locations []Location
}
