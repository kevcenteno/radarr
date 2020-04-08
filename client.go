package radarr

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

// HTTPClientInterface interface for the http.Client
type HTTPClientInterface interface {
	Get(url string) (resp *http.Response, err error)
	Do(req *http.Request) (*http.Response, error)
}

// New Create a Radarr client
// Optionnally specify an http.Client
func New(radarrURL, apiKey string, client HTTPClientInterface) (*Service, error) {
	valid, err := url.ParseRequestURI(radarrURL)
	if err != nil {
		return nil, errors.New("Please specify a valid URL")
	}

	// if client not specified, defaulting to default http client
	if client == nil {
		d := http.DefaultClient
		d.Transport = newTransport()
		d.Timeout = time.Second * 10
		client = d
	}

	s := &Service{client: client, url: valid.String(), apiKey: apiKey}
	s.Movies = newMovieService(s)
	s.SystemStatus = newSystemStatusService(s)
	s.Diskspace = newDiskspaceService(s)
	s.Command = newCommandService(s)
	s.History = newHistoryService(s)

	return s, nil
}

// Service containing all availables operations
type Service struct {
	client HTTPClientInterface
	url    string // Radarr base URL
	apiKey string

	// https://github.com/Radarr/Radarr/wiki/API:Calendar
	// https://github.com/Radarr/Radarr/wiki/API:Movie
	// https://github.com/Radarr/Radarr/wiki/API:Movie-Lookup
	Movies *MovieService

	// https://github.com/Radarr/Radarr/wiki/API:System-Status
	SystemStatus *SystemStatusService

	// https://github.com/Radarr/Radarr/wiki/API:Diskspace
	Diskspace *DiskspaceService

	// https://github.com/Radarr/Radarr/wiki/API:Command
	Command *CommandService

	// https://github.com/Radarr/Radarr/wiki/API:History
	History *HistoryService
}
