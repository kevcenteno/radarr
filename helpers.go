package radarr

import (
	"encoding/json"
	"net/http"
)

func parseRadarrResponse(response *http.Response) error {
	if response.StatusCode == http.StatusForbidden || response.StatusCode == http.StatusNotFound || response.StatusCode == http.StatusUnauthorized {
		e := Error{}
		e.Code = response.StatusCode
		e.Message = "Unknown"

		// Because Radarr response contains 'error' key or 'message' key. Parse it, and set to e.Message
		var body map[string]string
		err := json.NewDecoder(response.Body).Decode(&body)
		if err != nil {
			return err
		}
		if body["error"] != "" {
			e.Message = body["error"]
		} else if body["message"] != "" {
			e.Message = body["message"]
		}
		// Return final error
		return &e
	}
	return nil
}
