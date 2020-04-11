package radarr

// ImportMode can be used to override the default Copy for torrents with external preprocessing/transcoding/unrar.
type ImportMode int

const (
	// Move imported files instead of copy
	Move ImportMode = iota

	// Copy Or Hardlink depending on Radarr configuration
	Copy
)

// DownloadedMoviesScanOptions available options when using DownloadedMoviesScanCommand
type DownloadedMoviesScanOptions struct {
	Path             string     `json:"path"`
	DownloadClientID string     `json:"downloadClientId"`
	ImportMode       ImportMode `json:"importMode"`
}

var availableImportMode [2]string = [2]string{"Move", "Copy"}

func (i ImportMode) String() string {
	return availableImportMode[i]
}

type filter struct {
	Key   string      `json:"filterKey"`
	Value interface{} `json:"filterValue"`
}

// Here the list of all filters for movies
// If you have a better idea, please tell me
// I've been looking for every possible solution for seven hours... :'(
var availableFilter [7]filter = [7]filter{
	{Key: "monitored", Value: true},
	{Key: "monitored", Value: false},
	{Key: "all", Value: "all"},
	{Key: "status", Value: "available"},
	{Key: "status", Value: "released"},
	{Key: "status", Value: "inCinemas"},
	{Key: "status", Value: "announced"},
}

// Filter filtering options when using MissingMoviesSearch and CutOffUnmetMoviesSearchCommand
type Filter int

const (
	// FilterByMonitored filter movies by monitored ones
	FilterByMonitored Filter = iota

	// FilterByNonMonitored filter movies by non monitored ones
	FilterByNonMonitored

	// FilterAll return all movies without filters
	FilterAll

	// FilterByStatusAndAvailable return 'availables' movies
	FilterByStatusAndAvailable

	// FilterByStatusAndReleased return 'released' movies
	FilterByStatusAndReleased

	// FilterByStatusAndInCinemas return 'inCinemas' movies
	FilterByStatusAndInCinemas

	// FilterByStatusAndAnnounced return 'announced' movies
	FilterByStatusAndAnnounced
)

func (f Filter) get() *filter {
	return &availableFilter[f]
}
