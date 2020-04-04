package radarr

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func TestParseRadarrResponse(t *testing.T) {
	response := http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       http.NoBody,
	}

	// Parse a response 200
	err := parseRadarrResponse(&response)
	cases := []internal.TestCase{
		internal.TestCase{
			Title:    "Error should be nil",
			Expected: true,
			Got:      err == nil,
		},
	}

	// Test condition on statusCode with Unknown message
	status := [3]int{http.StatusNotFound, http.StatusForbidden, http.StatusUnauthorized}
	var expectedMessage string
	for _, s := range status {
		t.Run(http.StatusText(s), func(t *testing.T) {
			r := &http.Response{
				StatusCode: s,
				Body:       ioutil.NopCloser(bytes.NewBufferString("foo")),
			}
			err = parseRadarrResponse(r)
			if err == nil {
				t.Error("Error should not be nil")
			}
			expectedMessage = fmt.Sprintf("Radarr error: code %d, message '%s'", s, "Unknown")
			if err.Error() != expectedMessage {
				t.Errorf("Got'%s' want '%s'", err.Error(), expectedMessage)
			}

		})
	}

	// Test 'error' and 'message' key
	keys := map[string]map[string]string{
		"error": map[string]string{
			"response":         internal.DummyUnauthorizedResponse,
			"expected_message": "Unauthorized",
		},
		"message": map[string]string{
			"response":         internal.DummyNotFoundResponse,
			"expected_message": "NotFound",
		},
	}
	// For each keys, create a new response with respond responseBody["response"].
	// The expected message should be responseBody["expected_message"]
	for key, responseBody := range keys {
		response = http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(bytes.NewBufferString(responseBody["response"])),
		}
		err = parseRadarrResponse(&response)

		// Cast our error
		v, ok := err.(*Error)
		if !ok {
			t.Error("err should be an instance of Error")
		}
		cases = append(cases, []internal.TestCase{
			internal.TestCase{
				Title:    fmt.Sprintf("Key '%s': error message: %s", key, http.StatusText(http.StatusUnauthorized)),
				Expected: responseBody["expected_message"],
				Got:      v.Message,
			},
			internal.TestCase{
				Title:    fmt.Sprintf("Key '%s': message code: %s", key, http.StatusText(http.StatusUnauthorized)),
				Expected: http.StatusUnauthorized,
				Got:      v.Code,
			},
			internal.TestCase{
				Title:    fmt.Sprintf("Key '%s': message Error(): %s", key, http.StatusText(http.StatusUnauthorized)),
				Expected: fmt.Sprintf("Radarr error: code %d, message '%s'", http.StatusUnauthorized, responseBody["expected_message"]),
				Got:      v.Error(),
			},
		}...)
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			if c.Expected != c.Got {
				t.Errorf("Got '%v' want '%v'", c.Got, c.Expected)
			}
		})
	}
}
