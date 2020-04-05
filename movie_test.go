package radarr

import (
	"testing"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func TestGetMovie(t *testing.T) {
	// Init service
	service := newMovieService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
		apiKey: internal.DummyAPIKey,
	})

	// Search a non-existing movie
	movie, err := service.Get(123456789)
	cases := []internal.TestCase{
		internal.TestCase{
			Title:    "Error should be not nil",
			Expected: false,
			Got:      err == nil,
		},
		internal.TestCase{
			Title:    "Error message should contain NotFound",
			Expected: "Radarr error: code 404, message 'NotFound'",
			Got:      err.Error(),
		},
		internal.TestCase{
			Title:    "Movie should be nil",
			Expected: true,
			Got:      movie == nil,
		},
	}

	// Search an existing movie
	movie, err = service.Get(217)
	expectedTitle := "Frozen II"
	cases = append(cases, []internal.TestCase{
		internal.TestCase{
			Title:    "Error should be nil",
			Expected: nil,
			Got:      err,
		},
		internal.TestCase{
			Title:    "Movie should be not nil",
			Expected: false,
			Got:      movie == nil,
		},
		internal.TestCase{
			Title:    "Movie title should be correct",
			Expected: expectedTitle,
			Got:      movie.Title,
		},
	}...)

	// Bad API key
	service.s.apiKey = "foo"
	movie, err = service.Get(217)
	cases = append(cases, []internal.TestCase{
		internal.TestCase{
			Title:    "Movie should be nil because of bad API key",
			Expected: true,
			Got:      movie == nil,
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

func TestListMovies(t *testing.T) {
	// Init service
	service := newMovieService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
		apiKey: internal.DummyAPIKey,
	})

	movies, err := service.List()
	expectedTitle := "Frozen II"
	m := *movies
	cases := []internal.TestCase{
		internal.TestCase{
			Title:    "Movies count should be 2",
			Expected: 2,
			Got:      len(m),
		},
		internal.TestCase{
			Title:    "Error should be nil",
			Expected: true,
			Got:      err == nil,
		},
		internal.TestCase{
			Title:    "Title should match",
			Expected: expectedTitle,
			Got:      m[0].Title,
		},
		internal.TestCase{
			Title:    "Title should match",
			Expected: expectedTitle,
			Got:      m[1].Title,
		},
	}

	// Bad api key
	service.s.apiKey = "foo"
	movies, err = service.List()
	cases = append(cases, []internal.TestCase{
		internal.TestCase{
			Title:    "Error message should contain Unauthorized",
			Expected: "Radarr error: code 401, message 'Unauthorized'",
			Got:      err.Error(),
		},
		internal.TestCase{
			Title:    "Movies should be nil",
			Expected: true,
			Got:      movies == nil,
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
