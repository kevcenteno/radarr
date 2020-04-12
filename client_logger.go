package radarr

import (
	"net/http"
)

// Custom http transport to log request
type transport struct {
	transport http.RoundTripper
	apiKey    string
}

func newTransport(apiKey string) *transport {
	return &transport{
		apiKey:    apiKey,
		transport: http.DefaultTransport,
	}
}

// Custom RoundTrip method to insert custom headers on each requests
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "SkYNewZ-Go-http-client/1.1")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("X-Api-Key", t.apiKey)

	// Send request to the orignal transport
	resp, err := t.transport.RoundTrip(req)
	return resp, err
}
