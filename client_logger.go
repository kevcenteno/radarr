package radarr

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

const (
	envLog string = "LOG_LEVEL"
)

const logReqMsg = `API Request Details:
---[ REQUEST ]---------------------------------------
%s
-----------------------------------------------------
`

const logRespMsg = `API Response Details:
---[ RESPONSE ]--------------------------------------
%s
-----------------------------------------------------
`

// Custom http transport to log request
type transport struct {
	transport http.RoundTripper
}

func newTransport() *transport {
	return &transport{transport: http.DefaultTransport}
}

// Custom RoundTrip method to log every requests
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "SkYNewZ-Go-http-client/1.1")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Print request
	if isDebugLevel() {
		reqData, _ := httputil.DumpRequestOut(req, true)
		fmt.Printf(logReqMsg, reqData)
	}

	resp, err := t.transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Print response
	if isDebugLevel() {
		respData, _ := httputil.DumpResponse(resp, true)
		fmt.Printf(logRespMsg, respData)
	}

	// Follow request
	return resp, nil
}

func isDebugLevel() bool {
	return os.Getenv(envLog) == "DEBUG"
}
