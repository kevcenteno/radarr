package radarr

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Custom http transport to log request
type transport struct {
	transport http.RoundTripper
	apiKey    string
	verbose   bool
}

func newTransport(apiKey string, verbose bool) *transport {
	return &transport{
		apiKey:    apiKey,
		transport: http.DefaultTransport,
		verbose:   verbose,
	}
}

// Custom RoundTrip method to insert custom headers on each requests
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "SkYNewZ-Go-http-client/1.1")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("X-Api-Key", t.apiKey)

	if t.verbose {
		data, err := httputil.DumpRequest(req, true)
		if err != nil {
			return nil, err
		}
		fmt.Printf("[DEBUG] Request: %s\n", string(data))
	}

	// Send request to the orignal transport
	resp, err := t.transport.RoundTrip(req)

	if t.verbose {
		data, err := httputil.DumpResponse(resp, false)
		if err != nil {
			return nil, err
		}
		fmt.Printf("[DEBUG] Response: %s\n", string(data))
	}

	return resp, err
}
