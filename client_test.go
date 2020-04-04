package radarr

import (
	"net/http"
	"testing"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func TestServiceIsWellInstanciated(t *testing.T) {
	// Bad URL
	service, err := New("bad-url", internal.DummyAPIKey, internal.DummyHTTPClient)
	cases := []internal.TestCase{
		internal.TestCase{
			Title:    "Invalid URL error should be not nil",
			Expected: false,
			Got:      err == nil,
		},
		internal.TestCase{
			Title:    "Error message",
			Expected: "Please specify a valid URL",
			Got:      err.Error(),
		},
	}

	// Good service
	service, err = New(internal.DummyURL, internal.DummyAPIKey, internal.DummyHTTPClient)
	cases = append(cases, []internal.TestCase{
		internal.TestCase{
			Title:    "HTTP client service should be our testing service",
			Expected: true,
			Got:      service.client == internal.DummyHTTPClient,
		},
		internal.TestCase{
			Title:    "Error should be nil",
			Expected: nil,
			Got:      err,
		},
	}...)

	// Test different services
	cases = append(cases, []internal.TestCase{
		internal.TestCase{
			Title:    "Movies service should be not nil",
			Expected: false,
			Got:      service.Movies == nil,
		},
		internal.TestCase{
			Title:    "SystemStatus should be not nil",
			Expected: false,
			Got:      service.SystemStatus == nil,
		},
	}...)

	// Service with default http client
	service, _ = New(internal.DummyURL, internal.DummyAPIKey, nil)
	cases = append(cases, internal.TestCase{
		Title:    "HTTP client service should be the default http client",
		Expected: http.DefaultClient,
		Got:      service.client,
	})

	// Test attributes
	cases = append(cases, []internal.TestCase{
		internal.TestCase{
			Title:    "URL should be the same",
			Expected: internal.DummyURL,
			Got:      service.url,
		},
		internal.TestCase{
			Title:    "API key should be the same",
			Expected: internal.DummyAPIKey,
			Got:      service.apiKey,
		},
	}...)

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			if c.Expected != c.Got {
				t.Errorf("Got '%v' want '%v'", c.Got, c.Expected)
			}
		})
	}
}
