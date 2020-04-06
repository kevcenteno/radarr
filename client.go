package radarr

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// HTTPClientInterface interface for the http.Client
type HTTPClientInterface interface {
	Get(url string) (resp *http.Response, err error)
}

// New Create a Radarr client
// Optionnally specify an http.Client
func New(radarrURL, apiKey string, client HTTPClientInterface) (*Service, error) {
	valid, err := url.ParseRequestURI(radarrURL)
	if err != nil {
		return nil, fmt.Errorf("Please specify a valid URL")
	}

	// if client not specified, defaulting to default http client
	// TODO: test it
	if client == nil {
		d := http.DefaultClient
		d.Transport = newTransport()
		d.Timeout = time.Second * 10
		client = d
	}

	s := &Service{client: client, url: valid.String(), apiKey: apiKey}
	s.Movies = newMovieService(s)
	s.SystemStatus = newSystemStatusService(s)

	return s, nil
}

// Service containing all availables operations
type Service struct {
	client HTTPClientInterface
	url    string // Radarr base URL
	apiKey string

	Movies       *MovieService
	SystemStatus *SystemStatusService
}
