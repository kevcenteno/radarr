package radarr

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_parseRadarrResponse(t *testing.T) {
	tests := []struct {
		name        string
		response    *http.Response
		wantErr     bool
		wantMessage string
	}{
		{
			name: "No error",
			response: &http.Response{
				Status:     http.StatusText(http.StatusOK),
				StatusCode: http.StatusOK,
				Body:       http.NoBody,
			},
			wantErr: false,
		},
		{
			name: "Body does not contains 'error' or 'message' keys",
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"foo": "bar"}`)),
			},
			wantErr:     true,
			wantMessage: fmt.Sprintf("Radarr error: code %d, message 'Unable to read Radarr response body: foo=bar'", http.StatusBadRequest),
		},
		{
			name: "No body",
			response: &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       http.NoBody,
			},
			wantErr:     true,
			wantMessage: fmt.Sprintf("Radarr error: code %d, message 'Unknown'", http.StatusNotFound),
		},
		{
			name: "Not a JSON body",
			response: &http.Response{
				Status:     http.StatusText(http.StatusForbidden),
				StatusCode: http.StatusForbidden,
				Body:       ioutil.NopCloser(bytes.NewBufferString("foo")),
			},
			wantErr:     true,
			wantMessage: fmt.Sprintf("Radarr error: code %d, message 'foo'", http.StatusForbidden),
		},
	}

	// Test 'error' and 'message' key
	keys := map[string]map[string]string{
		"error": {
			"response":         `{"error": "Unauthorized"}`,
			"expected_message": "Unauthorized",
		},
		"message": {
			"response":         `{"message": "NotFound"}`,
			"expected_message": "NotFound",
		},
	}

	// For each keys, create a new response with respond responseBody["response"].
	// The expected message should be responseBody["expected_message"]
	for key, responseBody := range keys {
		response := &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(bytes.NewBufferString(responseBody["response"])),
		}

		tests = append(tests, struct {
			name        string
			response    *http.Response
			wantErr     bool
			wantMessage string
		}{
			name:        fmt.Sprintf("Key '%s': message Error(): %s", key, http.StatusText(http.StatusUnauthorized)),
			response:    response,
			wantErr:     true,
			wantMessage: fmt.Sprintf("Radarr error: code %d, message '%s'", http.StatusUnauthorized, responseBody["expected_message"]),
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parseRadarrResponse(tt.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRadarrResponse() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err.Error() != tt.wantMessage {
				t.Errorf("err.Error() error = %s, wantMessage %s", err.Error(), tt.wantMessage)
			}
		})
	}
}
