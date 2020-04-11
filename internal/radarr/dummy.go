// Package radarr here only exist for testing
package radarr

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// DummyMovieResponse describe a dymmy movie
var DummyMovieResponse string = `
{
  "title": "Frozen II",
  "alternativeTitles": [
    {
      "sourceType": "tmdb",
      "movieId": 217,
      "title": "Frozen 2",
      "sourceId": 330457,
      "votes": 0,
      "voteCount": 0,
      "language": {
        "id": 1,
        "name": "English"
      },
      "id": 461
    }
  ],
  "secondaryYearSourceId": 0,
  "sortTitle": "frozen ii",
  "sizeOnDisk": 4099483594,
  "status": "released",
  "overview": "Elsa, Anna, Kristoff and Olaf head far into the forest to learn the truth about an ancient mystery of their kingdom.",
  "inCinemas": "2019-11-19T23:00:00Z",
  "physicalRelease": "2020-02-11T00:00:00Z",
  "images": [
    {
      "coverType": "poster",
      "url": "/MediaCover/217/poster.jpg?lastWrite=637214530603577317"
    },
    {
      "coverType": "fanart",
      "url": "/MediaCover/217/fanart.jpg?lastWrite=637202450497927734"
    }
  ],
  "website": "https://movies.disney.com/frozen-2",
  "downloaded": true,
  "year": 2019,
  "hasFile": true,
  "youTubeTrailerId": "Zi4LMpSDccc",
  "studio": "Walt Disney Animation Studios",
  "path": "/movies/Frozen II (2019)",
  "profileId": 3,
  "monitored": true,
  "minimumAvailability": "inCinemas",
  "isAvailable": true,
  "folderName": "/movies/Frozen II (2019)",
  "runtime": 104,
  "lastInfoSync": "2020-04-03T19:39:49.6265379Z",
  "cleanTitle": "frozenii",
  "imdbId": "tt4520988",
  "tmdbId": 330457,
  "titleSlug": "frozen-ii-330457",
  "genres": ["Animation", "Family", "Adventure"],
  "tags": [1],
  "added": "2020-03-15T15:39:15.8796553Z",
  "ratings": {
    "votes": 3481,
    "value": 7.1
  },
  "movieFile": {
    "movieId": 0,
    "relativePath": "Frozen.2.2019.MULTi.1080p.WEB.x264.EXTREME.mkv",
    "size": 4099483594,
    "dateAdded": "2020-03-15T16:18:06.9156804Z",
    "sceneName": "Frozen.2.2019.MULTi.1080p.WEB.x264.EXTREME",
    "quality": {
      "quality": {
        "id": 3,
        "name": "WEBDL-1080p",
        "source": "webdl",
        "resolution": 1080,
        "modifier": "none"
      },
      "revision": {
        "version": 1,
        "real": 0,
        "isRepack": false
      }
    },
    "edition": "",
    "mediaInfo": {
      "containerFormat": "Matroska",
      "videoFormat": "AVC",
      "videoCodecID": "V_MPEG4/ISO/AVC",
      "videoProfile": "High@L4",
      "videoCodecLibrary": "",
      "videoBitrate": 4526004,
      "videoBitDepth": 8,
      "videoMultiViewCount": 0,
      "videoColourPrimaries": "BT.709",
      "videoTransferCharacteristics": "BT.709",
      "width": 1920,
      "height": 804,
      "audioFormat": "AC-3",
      "audioCodecID": "A_AC3",
      "audioCodecLibrary": "",
      "audioAdditionalFeatures": "",
      "audioBitrate": 384000,
      "runTime": "01:43:12.3530000",
      "audioStreamCount": 2,
      "audioChannels": 6,
      "audioChannelPositions": "3/2/0.1",
      "audioChannelPositionsText": "Front: L C R, Side: L R, LFE",
      "audioProfile": "",
      "videoFps": 23.976,
      "audioLanguages": "French / English",
      "subtitles": "French",
      "scanType": "Progressive",
      "schemaRevision": 5
    },
    "id": 197
  },
  "qualityProfileId": 3,
  "id": 217
}`

var dummySystemStatusResponse string = `
{
  "version": "3.0.0.2741",
  "buildTime": "2020-03-23T16:23:16Z",
  "isDebug": false,
  "isProduction": true,
  "isAdmin": false,
  "isUserInteractive": false,
  "startupPath": "/opt/radarr",
  "appData": "/config",
  "osName": "ubuntu",
  "osVersion": "20.04",
  "isNetCore": true,
  "isMono": false,
  "isLinux": true,
  "isOsx": false,
  "isWindows": false,
  "branch": "develop",
  "authentication": "forms",
  "sqliteVersion": "3.31.1",
  "migrationVersion": 169,
  "urlBase": "",
  "runtimeVersion": "3.1.2",
  "runtimeName": "netCore"
}`

var dummyDiskspaceResponse string = `
[{
  "path": "/",
  "label": "/",
  "freeSpace": 11059175424,
  "totalSpace": 15614754816
},
{
  "path": "/home",
  "label": "/home",
  "freeSpace": 88775757824,
  "totalSpace": 98325770240
}
]`

// DummyHistoryResponse describe /history response
var DummyHistoryResponse string = `
{
	"page": 1,
	"pageSize": 1,
	"sortKey": "date",
	"sortDirection": "descending",
	"totalRecords": 131,
	"records": [{
		"movieId": 194,
		"sourceTitle": "/movies/Ford v Ferrari (2019)/Ford.v.Ferrari.2019.MULTi.2160p.UHD.BluRay.REMUX.HEVC-BEO.mkv",
		"quality": {
			"quality": {
				"id": 31,
				"name": "Remux-2160p",
				"source": "bluray",
				"resolution": 2160,
				"modifier": "remux"
			},
			"revision": {
				"version": 1,
				"real": 0,
				"isRepack": false
			}
		},
		"qualityCutoffNotMet": false,
		"date": "2020-04-05T19:43:55.5957884Z",
		"eventType": "movieFileDeleted",
		"data": {
			"reason": "MissingFromDisk"
		},
		"movie": {
			"title": "Le Mans 66",
			"alternativeTitles": [],
			"secondaryYearSourceId": 0,
			"sortTitle": "le mans 66",
			"sizeOnDisk": 0,
			"status": "released",
			"overview": "Relate l’histoire vraie qui a conduit l’ingénieur automobile visionnaire américain Caroll Shelby à faire équipe avec le pilote de course britannique surdoué Ken Miles. Bravant l’ordre établi, défiant les lois de la physique et luttant contre leurs propres démons, les deux hommes n’avaient qu’un seul but: construire pour le compte de Ford Motor Company un bolide révolutionnaire capable de renverser la suprématie de l’écurie d’Enzo Ferrari sur le mythique circuit des 24 heures du Mans en 1966…",
			"inCinemas": "2019-11-12T23:00:00Z",
			"physicalRelease": "2020-01-28T00:00:00Z",
			"images": [{
					"coverType": "poster",
					"url": "http://image.tmdb.org/t/p/original/8yyRujXGSNCa3yrM3qoLZXUW3WY.jpg"
				},
				{
					"coverType": "fanart",
					"url": "http://image.tmdb.org/t/p/original/n3UanIvmnBlH531pykuzNs4LbH6.jpg"
				}
			],
			"website": "https://www.foxmovies.com/movies/ford-v-ferrari",
			"downloaded": false,
			"year": 2019,
			"hasFile": false,
			"youTubeTrailerId": "EVZbiA81v7w",
			"studio": "20th Century Fox",
			"path": "/movies/Ford v Ferrari (2019)",
			"profileId": 5,
			"monitored": false,
			"minimumAvailability": "released",
			"isAvailable": true,
			"folderName": "/movies/Ford v Ferrari (2019)",
			"runtime": 152,
			"lastInfoSync": "2020-04-07T19:51:24.7882218Z",
			"cleanTitle": "lemans66",
			"imdbId": "tt1950186",
			"tmdbId": 359724,
			"titleSlug": "le-mans-66-359724",
			"genres": [
				"Drame",
				"Action"
			],
			"tags": [],
			"added": "2019-12-15T16:26:32.3913355Z",
			"ratings": {
				"votes": 2228,
				"value": 7.8
			},
			"qualityProfileId": 5,
			"id": 194
		},
		"id": 138
	}]
}`

var dummyMoviesResponse string = fmt.Sprintf("[%s, %s]", DummyMovieResponse, DummyMovieResponse)
var dummyUpcomingWithBothFilterResponse = fmt.Sprintf("[%s]", DummyMovieResponse)

var dummyNotFoundResponse *http.Response = &http.Response{
	StatusCode: http.StatusNotFound,
	Status:     http.StatusText(http.StatusNotFound),
	Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "NotFound"}`)),
}

var dummyEmptyListResponse string = `[]`

var dummyStartDate string = "2019-11-19T23:00:00Z"
var dummyEndDate string = "2019-11-20T23:00:00Z"

var dummyGenericResponse = &http.Response{
	StatusCode: http.StatusOK,
	Status:     http.StatusText(http.StatusOK),
	Body:       ioutil.NopCloser(bytes.NewBufferString(`{"foo": "bar"}`)),
}

var (
	// DummyHTTPClient mocked http client
	DummyHTTPClient *HTTPClient

	// DummyURL dummy Radarr instance URL
	DummyURL string = "https://radarr.dummy"

	// DummyAPIKey dummy Radarr API keys
	DummyAPIKey string = "dummy-api-key"

	// ParseDummyURL parsed dummy URL
	ParseDummyURL, _ = url.Parse(DummyURL)
)

type mockedTransport1 struct {
	http.RoundTripper
	MockedResponse *http.Response
}
type mockedTransport2 struct {
	http.RoundTripper
	MockedResponse *http.Response
}

// MockedTransports mocked http.Client transport
type MockedTransports struct {
	MockedTransport1 *mockedTransport1
	MockedTransport2 *mockedTransport2
}

// NewMockedTransports MockedTransports constructor
func NewMockedTransports() *MockedTransports {
	return &MockedTransports{
		MockedTransport1: &mockedTransport1{MockedResponse: dummyGenericResponse},
		MockedTransport2: &mockedTransport2{MockedResponse: dummyGenericResponse},
	}
}

func (r *mockedTransport1) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.MockedResponse, nil
}

func (*mockedTransport2) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, errors.New("foo")
}

// HTTPClient implements HTTPClientInterface
type HTTPClient struct{}

func init() {
	// Create a mock http client
	DummyHTTPClient = &HTTPClient{}
}

// Do mocked http client Do function
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	// Query params
	params := req.URL.Query()

	if req.Method == http.MethodGet {
		// GET /history
		if strings.Contains(req.URL.String(), fmt.Sprintf("%s/api%s", DummyURL, "/history")) {
			page := params.Get("page")
			pageSize := params.Get("pageSize")

			// Return one record on page 1
			if page == "1" && pageSize == "50" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(DummyHistoryResponse)),
				}, nil
			}

			// Return bad JSON for page 3
			if page == "3" && pageSize == "50" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString("foo")),
				}, nil
			}

			// Return error for page 4
			if page == "4" && pageSize == "50" {
				return nil, errors.New("Oooops")
			}
		}

		// GET /calendar
		if strings.Contains(req.URL.String(), fmt.Sprintf("%s/api%s", DummyURL, "/calendar")) {
			start := params.Get("start")
			end := params.Get("end")

			// Upcoming movies without filters. Return 0 movies
			if start == "" && end == "" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(dummyEmptyListResponse)),
				}, nil
			}

			// Upcoming movies with start filter. Returns 2 movies
			if start == dummyStartDate && end == "" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(dummyMoviesResponse)),
				}, nil
			}

			// Upcoming movies with end filter. Returns 0 movies
			if start == "" && end == dummyEndDate {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(dummyEmptyListResponse)),
				}, nil
			}

			// Upcoming movies with start filter and end filter. Return 1 movies
			if start == dummyStartDate && end == dummyEndDate {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(dummyUpcomingWithBothFilterResponse)),
				}, nil
			}
		}

		// Else, return 404
		return dummyNotFoundResponse, nil
	}

	if req.Method == http.MethodDelete {
		// Delete movie
		if strings.Contains(req.URL.String(), fmt.Sprintf("%s/api%s/%d", DummyURL, "/movie", 217)) {
			deleteFiles := params.Get("deleteFiles")
			addExclusion := params.Get("addExclusion")

			// without parameters
			if addExclusion == "" && deleteFiles == "" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
				}, nil
			}

			// With all false
			if addExclusion == "false" && deleteFiles == "false" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
				}, nil
			}

			// With all true
			if addExclusion == "true" && deleteFiles == "true" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
				}, nil
			}

			// With addExclusion=true and deleteFiles=false
			if addExclusion == "true" && deleteFiles == "false" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
				}, nil
			}

			// With addExclusion=false and deleteFiles=true
			if addExclusion == "false" && deleteFiles == "true" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
					Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
				}, nil
			}
		}

		return dummyNotFoundResponse, nil
	}

	// Bad method, return error
	return nil, &url.Error{
		Op:  req.Method,
		URL: req.URL.String(),
		Err: errors.New("Ooops"),
	}
}

// Get Mock .Get() http client method
func (c *HTTPClient) Get(targetURL string) (resp *http.Response, err error) {
	switch targetURL {
	case fmt.Sprintf("%s/api%s/%d", DummyURL, "/movie", 217):
		// Get one movie
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     http.StatusText(http.StatusOK),
			Body:       ioutil.NopCloser(bytes.NewBufferString(DummyMovieResponse)),
		}, nil

	case fmt.Sprintf("%s/api%s", DummyURL, "/movie"):
		// List of movies
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     http.StatusText(http.StatusOK),
			Body:       ioutil.NopCloser(bytes.NewBufferString(dummyMoviesResponse)),
		}, nil

	case fmt.Sprintf("%s/api%s", DummyURL, "/system/status"):
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     http.StatusText(http.StatusOK),
			Body:       ioutil.NopCloser(bytes.NewBufferString(dummySystemStatusResponse)),
		}, nil

	case fmt.Sprintf("%s/api%s", DummyURL, "/diskspace"):
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     http.StatusText(http.StatusOK),
			Body:       ioutil.NopCloser(bytes.NewBufferString(dummyDiskspaceResponse)),
		}, nil

	default:
		// Defaulting to 404
		return dummyNotFoundResponse, nil
	}
}
