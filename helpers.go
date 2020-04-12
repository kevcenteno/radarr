package radarr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func parseRadarrResponse(response *http.Response) error {
	// If ok, send no error
	if response.StatusCode == http.StatusOK {
		return nil
	}

	e := Error{}
	e.Code = response.StatusCode
	e.Message = "Unknown"

	// Read response body as plain text
	data, _ := ioutil.ReadAll(response.Body)

	// Because Radarr response contains 'error' key or 'message' key. Parse it, and set to e.Message
	var body map[string]string
	err := json.Unmarshal(data, &body)

	if err != nil {
		switch {
		case len(data) == 0:
			// No body
			return &e
		default:
			// Body as plain text
			e.Message = string(data)
			return &e
		}
	}

	// If JSON decoding do not fail
	switch {
	case body["error"] != "":
		e.Message = body["error"]
		return &e
	case body["message"] != "":
		e.Message = body["message"]
		return &e
	default:
		var m string
		for key, value := range body {
			m += fmt.Sprintf("%s=%s", key, value)
		}

		e.Message = fmt.Sprintf("Unable to read Radarr response body: %s", m)
		return &e
	}
}
