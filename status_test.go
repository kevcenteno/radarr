package radarr

import (
	"testing"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func TestGet(t *testing.T) {
	service := newSystemStatusService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
		apiKey: internal.DummyAPIKey,
	})
	dummySystemStatusResponse, err := service.Get()
	cases := []internal.TestCase{
		internal.TestCase{
			Title:    "Error should be nil",
			Expected: true,
			Got:      err == nil,
		},
		internal.TestCase{
			Title:    "Response should be not nil",
			Expected: false,
			Got:      dummySystemStatusResponse == nil,
		},
		internal.TestCase{
			Title:    "Version should be same",
			Expected: "3.0.0.2741",
			Got:      dummySystemStatusResponse.Version,
		},
	}

	// Bad API key
	service.s.apiKey = "foo"
	dummySystemStatusResponse, err = service.Get()
	cases = append(cases, []internal.TestCase{
		internal.TestCase{
			Title:    "Response should be nil because of bad API key",
			Expected: true,
			Got:      dummySystemStatusResponse == nil,
		},
		internal.TestCase{
			Title:    "Error should be not nil",
			Expected: false,
			Got:      err == nil,
		},
		internal.TestCase{
			Title:    "Error message should contain Unauthorized",
			Expected: "Radarr error: code 401, message 'Unauthorized'",
			Got:      err.Error(),
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
