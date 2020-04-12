package radarr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

// LookupService contains Radarr movie lookup operations
type LookupService struct {
	s *Service
}

func newLookupService(s *Service) *LookupService {
	return &LookupService{s}
}

// Plain simply search movies matching given term
// https://github.com/Radarr/Radarr/wiki/API:Movie-Lookup#search-by-term
func (l *LookupService) Plain(term string) (Movies, error) {
	if term == "" {
		return nil, errors.New("Please specify a valid term")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api%s", l.s.url, lookupURI), nil)
	if err != nil {
		return nil, err
	}

	// Append query
	params := req.URL.Query()
	params.Add("term", term)
	req.URL.RawQuery = params.Encode()

	resp, err := l.s.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = parseRadarrResponse(resp)
	if err != nil {
		return nil, err
	}

	var movies Movies
	err = json.NewDecoder(resp.Body).Decode(&movies)
	if err != nil {
		return nil, err
	}

	_ = resp.Body.Close()
	return movies, nil
}

// Tmdb Search by The Movie Database ID
// https://github.com/Radarr/Radarr/wiki/API:Movie-Lookup#search-by-the-movie-database-id
func (l *LookupService) Tmdb(TMDBID int) (*Movie, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api%s", l.s.url, lookupTMDBURI), nil)
	if err != nil {
		return nil, err
	}

	// Append query
	params := req.URL.Query()
	params.Add("tmdbId", strconv.Itoa(TMDBID))
	req.URL.RawQuery = params.Encode()

	resp, err := l.s.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = parseRadarrResponse(resp)
	if err != nil {
		return nil, err
	}

	var movie Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return nil, err
	}

	_ = resp.Body.Close()
	return &movie, nil
}

// Imdb Search using IMDB id
// https://github.com/Radarr/Radarr/wiki/API:Movie-Lookup#search-using-imdb-id
func (l *LookupService) Imdb(IMDBID string) (*Movie, error) {
	var re = regexp.MustCompile(`(?i)^tt\d*$`)
	if !re.MatchString(IMDBID) {
		return nil, fmt.Errorf("Please specify a valid IMDB id. Pattern %s", re.String())
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api%s", l.s.url, lookupIMDBURI), nil)
	if err != nil {
		return nil, err
	}

	// Append query
	parms := req.URL.Query()
	parms.Add("imdbId", IMDBID)
	req.URL.RawQuery = parms.Encode()

	resp, err := l.s.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = parseRadarrResponse(resp)
	if err != nil {
		return nil, err
	}

	var movie Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return nil, err
	}

	_ = resp.Body.Close()
	return &movie, nil
}
