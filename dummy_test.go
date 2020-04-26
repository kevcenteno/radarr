package radarr

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

type hTTPClient struct{}

var (
	dummyHTTPClient *hTTPClient = &hTTPClient{}
	dummyURL        string      = "https://radarr.dummy"
	dummyAPIKey     string      = "dummy-api-key"
)

// Read the given .json file
func helperLoadBytes(name string) []byte {
	path := filepath.Join("testdata", name+".json")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}

func dummyMoviesResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("multiple_movies"))),
	}
}

func dummyMovieResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("movie"))),
	}
}

func dummyUpcomingWithBothFilterResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("one_movie_in_array"))),
	}
}

func dummyEmptyListResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBufferString(`[]`))}
}

func dummyNotFoundResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     http.StatusText(http.StatusNotFound),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("not_found"))),
	}
}

func dummyHistoryResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("history"))),
	}
}

func dummyInvalidJSONResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBufferString("foo"))}
}

func dummyEmptyJSONObject() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`))}
}

func invalidIMDBIDResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Status:     http.StatusText(http.StatusInternalServerError),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("invalid_imdb_id"))),
	}
}

func invalidTMDBIDResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Status:     http.StatusText(http.StatusInternalServerError),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("invalid_tmdb_id"))),
	}
}

func dummySystemStatusResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("status"))),
	}
}

func dummyDiskspaceResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("diskspace"))),
	}
}

func dummyExcludedMovies() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(bytes.NewBuffer(helperLoadBytes("excluded_movie"))),
	}
}

var dummyStartDate string = "2019-11-19T23:00:00Z"
var dummyEndDate string = "2019-11-20T23:00:00Z"

// Do mocked http client Do function
func (c *hTTPClient) Do(req *http.Request) (*http.Response, error) {
	// Query params
	params := req.URL.Query()

	if req.Method == http.MethodGet {
		// GET /history
		if strings.Contains(req.URL.String(), fmt.Sprintf("%s/api%s", dummyURL, "/history")) {
			page := params.Get("page")
			pageSize := params.Get("pageSize")

			// Return one record on page 1
			if page == "1" && pageSize == "50" {
				return dummyHistoryResponse(), nil
			}

			// Return bad JSON for page 3
			if page == "3" && pageSize == "50" {
				return dummyInvalidJSONResponse(), nil
			}

			// Return error for page 4
			if page == "4" && pageSize == "50" {
				return nil, errors.New("Oooops")
			}
		}

		// GET /calendar
		if strings.Contains(req.URL.String(), fmt.Sprintf("%s/api%s", dummyURL, "/calendar")) {
			start := params.Get("start")
			end := params.Get("end")

			// Upcoming movies without filters. Return 0 movies
			if start == "" && end == "" {
				return dummyEmptyListResponse(), nil
			}

			// Upcoming movies with start filter. Returns 2 movies
			if start == dummyStartDate && end == "" {
				return dummyMoviesResponse(), nil
			}

			// Upcoming movies with end filter. Returns 0 movies
			if start == "" && end == dummyEndDate {
				return dummyEmptyListResponse(), nil
			}

			// Upcoming movies with start filter and end filter. Return 1 movies
			if start == dummyStartDate && end == dummyEndDate {
				return dummyUpcomingWithBothFilterResponse(), nil
			}
		}

		// GET /movie/lookup/imdb
		if strings.Contains(req.URL.String(), fmt.Sprintf("%s/api%s", dummyURL, "/movie/lookup/imdb")) {
			imdbID := params.Get("imdbId")
			switch imdbID {
			case "tt3778644":
				return dummyMovieResponse(), nil
			case "1":
				return invalidIMDBIDResponse(), nil
			}
		}

		// GET /movie/lookup/tmdb
		if strings.Contains(req.URL.String(), fmt.Sprintf("%s/api%s", dummyURL, "/movie/lookup/tmdb")) {
			tmdbID := params.Get("tmdbId")
			switch tmdbID {
			case "348350":
				return dummyMovieResponse(), nil
			case "1":
				return invalidTMDBIDResponse(), nil
			}
		}

		// GET /movie/lookup
		if strings.Contains(req.URL.String(), fmt.Sprintf("%s/api%s", dummyURL, "/movie/lookup")) {
			term := params.Get("term")
			switch term {
			case "star wars":
				return dummyMoviesResponse(), nil
			case "this film does not exist either":
				return dummyEmptyListResponse(), nil
			}
		}

		// Else, return 404
		return dummyNotFoundResponse(), nil
	}

	if req.Method == http.MethodDelete {
		// Delete movie
		if strings.Contains(req.URL.String(), fmt.Sprintf("%s/api%s/%d", dummyURL, "/movie", 217)) {
			deleteFiles := params.Get("deleteFiles")
			addExclusion := params.Get("addExclusion")

			// without parameters
			if addExclusion == "" && deleteFiles == "" {
				return dummyEmptyJSONObject(), nil
			}

			// With all false
			if addExclusion == "false" && deleteFiles == "false" {
				return dummyEmptyJSONObject(), nil
			}

			// With all true
			if addExclusion == "true" && deleteFiles == "true" {
				return dummyEmptyJSONObject(), nil
			}

			// With addExclusion=true and deleteFiles=false
			if addExclusion == "true" && deleteFiles == "false" {
				return dummyEmptyJSONObject(), nil
			}

			// With addExclusion=false and deleteFiles=true
			if addExclusion == "false" && deleteFiles == "true" {
				return dummyEmptyJSONObject(), nil
			}
		}

		return dummyNotFoundResponse(), nil
	}

	// Bad method, return error
	return nil, &url.Error{
		Op:  req.Method,
		URL: req.URL.String(),
		Err: errors.New("Ooops"),
	}
}

// Get Mock .Get() http client method
func (c *hTTPClient) Get(targetURL string) (resp *http.Response, err error) {
	switch targetURL {
	case fmt.Sprintf("%s/api%s/%d", dummyURL, "/movie", 217):
		// Get one movie
		return dummyMovieResponse(), nil

	case fmt.Sprintf("%s/api%s", dummyURL, "/movie"):
		// List of movies
		return dummyMoviesResponse(), nil

	case fmt.Sprintf("%s/api%s", dummyURL, "/system/status"):
		return dummySystemStatusResponse(), nil

	case fmt.Sprintf("%s/api%s", dummyURL, "/diskspace"):
		return dummyDiskspaceResponse(), nil

	case fmt.Sprintf("%s/api%s", dummyURL, "/exclusions"):
		return dummyExcludedMovies(), nil

	default:
		// Defaulting to 404
		return dummyNotFoundResponse(), nil
	}
}
