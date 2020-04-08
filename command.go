package radarr

import (
	"time"
)

// Command describe a generic command
type Command struct {
	Name string `json:"name"`
	Body struct {
		SendUpdatesToClient bool   `json:"sendUpdatesToClient"`
		UpdateScheduledTask bool   `json:"updateScheduledTask"`
		CompletionMessage   string `json:"completionMessage"`
		Name                string `json:"name"`
		Trigger             string `json:"trigger"`
	} `json:"body"`
	Priority            string    `json:"priority"`
	Status              string    `json:"status"`
	Queued              time.Time `json:"queued"`
	Trigger             string    `json:"trigger"`
	State               string    `json:"state"`
	Manual              bool      `json:"manual"`
	StartedOn           time.Time `json:"startedOn"`
	StateChangeTime     time.Time `json:"stateChangeTime"`
	SendUpdatesToClient bool      `json:"sendUpdatesToClient"`
	UpdateScheduledTask bool      `json:"updateScheduledTask"`
	ID                  int       `json:"id"`
}

type Tasks []Command

// CommandService not usable for now
// contains Radarr commands operations
type CommandService struct {
	s *Service
}

func newCommandService(s *Service) *CommandService {
	return &CommandService{s}
}

// Status Queries the status of a previously started command mathing given unique ID
func (c *CommandService) Status(commandID string) *Command {
	return nil
}

// StatusAll Queries the status of all currently started commands.
func (c *CommandService) StatusAll() *Tasks {
	return nil
}

// RefreshMovie Refresh movie information from TMDb and rescan disk
func (c *CommandService) RefreshMovie(movieID ...int) *Command {
	// name := "RefreshMovie"
	return nil
}

// RescanMovie Rescan disk for movies
func (c *CommandService) RescanMovie(movieID ...int) *Command {
	// name := "RescanMovie"
	return nil
}

// MoviesSearch Search for one or more movies
func (c *CommandService) MoviesSearch(movieIDs ...[]int) *Command {
	// name := "MoviesSearch"
	return nil
}

// DownloadedMoviesScan Instruct Radarr to scan the DroneFactoryFolder or a folder defined by the path variable.
// Each file and folder in the DroneFactoryFolder is interpreted as separate download.
// But a folder specified by the path variable is assumed to be a single download (job) and the folder name should be the release name.
// The downloadClientId can be used to support this API endpoint in conjunction with Completed Download Handling, so Radarr knows that a particular download has already been imported.
func (c *CommandService) DownloadedMoviesScan(opts ...*DownloadedMoviesScanOptions) *Command {
	// name := "DownloadedMoviesScan"
	return nil
}

// RssSync Instruct Radarr to perform an RSS sync with all enabled indexers
func (c *CommandService) RssSync() *Command {
	// name := "RssSync"
	return nil
}

// RenameFiles Instruct Radarr to rename the list of files provided.
func (c *CommandService) RenameFiles(files ...[]int) *Command {
	// name := "RenameFiles"
	return nil
}

// RenameMovie Instruct Radarr to rename all files in the provided movies.
func (c *CommandService) RenameMovie(movieIDs ...[]int) *Command {
	// name := "RenameMovie"
	return nil
}

// CutOffUnmetMoviesSearch Instructs Radarr to search all cutoff unmet movies (Take care, since it could go over your indexers api limits!)
func (c *CommandService) CutOffUnmetMoviesSearch(filter *Filter) *Command {
	// name := "CutOffUnmetMoviesSearch"
	return nil
}

// NetImportSync Instructs Radarr to search all lists for movies not yet added to Radarr.
func (c *CommandService) NetImportSync() *Command {
	// name := "NetImportSync"
	return nil
}

// MissingMoviesSearch Instructs Radarr to search all missing movies.
// This functionality is similar to what CouchPotato does and runs a backlog search for all your missing movies.
// For example You can use this api with curl and crontab to instruct Radarr to run a backlog search on 1 AM everyday.
func (c *CommandService) MissingMoviesSearch(filter *Filter) *Command {
	// name := "MissingMoviesSearch"
	return nil
}
