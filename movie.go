package radarr

import (
	"encoding/json"
	"fmt"
	"time"
)

// Movie Radarr movie
type Movie struct {
	Title             string `json:"title"`
	AlternativeTitles []struct {
		SourceType string `json:"sourceType"`
		MovieID    int    `json:"movieId"`
		Title      string `json:"title"`
		SourceID   int    `json:"sourceId"`
		Votes      int    `json:"votes"`
		VoteCount  int    `json:"voteCount"`
		Language   struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"language"`
		ID int `json:"id"`
	} `json:"alternativeTitles"`
	SecondaryYearSourceID int       `json:"secondaryYearSourceId"`
	SortTitle             string    `json:"sortTitle"`
	SizeOnDisk            int64     `json:"sizeOnDisk"`
	Status                string    `json:"status"`
	Overview              string    `json:"overview"`
	InCinemas             time.Time `json:"inCinemas"`
	PhysicalRelease       time.Time `json:"physicalRelease"`
	Images                []struct {
		CoverType string `json:"coverType"`
		URL       string `json:"url"`
	} `json:"images"`
	Website             string    `json:"website"`
	Downloaded          bool      `json:"downloaded"`
	Year                int       `json:"year"`
	HasFile             bool      `json:"hasFile"`
	YouTubeTrailerID    string    `json:"youTubeTrailerId"`
	Studio              string    `json:"studio"`
	Path                string    `json:"path"`
	ProfileID           int       `json:"profileId"`
	Monitored           bool      `json:"monitored"`
	MinimumAvailability string    `json:"minimumAvailability"`
	IsAvailable         bool      `json:"isAvailable"`
	FolderName          string    `json:"folderName"`
	Runtime             int       `json:"runtime"`
	LastInfoSync        time.Time `json:"lastInfoSync"`
	CleanTitle          string    `json:"cleanTitle"`
	ImdbID              string    `json:"imdbId"`
	TmdbID              int       `json:"tmdbId"`
	TitleSlug           string    `json:"titleSlug"`
	Genres              []string  `json:"genres"`
	Tags                []int     `json:"tags"`
	Added               time.Time `json:"added"`
	Ratings             struct {
		Votes int     `json:"votes"`
		Value float64 `json:"value"`
	} `json:"ratings"`
	MovieFile struct {
		MovieID      int       `json:"movieId"`
		RelativePath string    `json:"relativePath"`
		Size         int64     `json:"size"`
		DateAdded    time.Time `json:"dateAdded"`
		SceneName    string    `json:"sceneName"`
		Quality      struct {
			Quality struct {
				ID         int    `json:"id"`
				Name       string `json:"name"`
				Source     string `json:"source"`
				Resolution int    `json:"resolution"`
				Modifier   string `json:"modifier"`
			} `json:"quality"`
			Revision struct {
				Version  int  `json:"version"`
				Real     int  `json:"real"`
				IsRepack bool `json:"isRepack"`
			} `json:"revision"`
		} `json:"quality"`
		Edition   string `json:"edition"`
		MediaInfo struct {
			ContainerFormat              string  `json:"containerFormat"`
			VideoFormat                  string  `json:"videoFormat"`
			VideoCodecID                 string  `json:"videoCodecID"`
			VideoProfile                 string  `json:"videoProfile"`
			VideoCodecLibrary            string  `json:"videoCodecLibrary"`
			VideoBitrate                 int     `json:"videoBitrate"`
			VideoBitDepth                int     `json:"videoBitDepth"`
			VideoMultiViewCount          int     `json:"videoMultiViewCount"`
			VideoColourPrimaries         string  `json:"videoColourPrimaries"`
			VideoTransferCharacteristics string  `json:"videoTransferCharacteristics"`
			Width                        int     `json:"width"`
			Height                       int     `json:"height"`
			AudioFormat                  string  `json:"audioFormat"`
			AudioCodecID                 string  `json:"audioCodecID"`
			AudioCodecLibrary            string  `json:"audioCodecLibrary"`
			AudioAdditionalFeatures      string  `json:"audioAdditionalFeatures"`
			AudioBitrate                 int     `json:"audioBitrate"`
			RunTime                      string  `json:"runTime"`
			AudioStreamCount             int     `json:"audioStreamCount"`
			AudioChannels                int     `json:"audioChannels"`
			AudioChannelPositions        string  `json:"audioChannelPositions"`
			AudioChannelPositionsText    string  `json:"audioChannelPositionsText"`
			AudioProfile                 string  `json:"audioProfile"`
			VideoFps                     float64 `json:"videoFps"`
			AudioLanguages               string  `json:"audioLanguages"`
			Subtitles                    string  `json:"subtitles"`
			ScanType                     string  `json:"scanType"`
			SchemaRevision               int     `json:"schemaRevision"`
		} `json:"mediaInfo"`
		ID int `json:"id"`
	} `json:"movieFile"`
	QualityProfileID int `json:"qualityProfileId"`
	ID               int `json:"id"`
}

// Movies multiple Radarr movies
type Movies []Movie

// MovieService contains Radarr movies operations
type MovieService struct {
	s *Service
}

func newMovieService(s *Service) *MovieService {
	return &MovieService{
		s: s,
	}
}

// GetMovie https://github.com/Radarr/Radarr/wiki/API:Movie#getid
func (m *MovieService) GetMovie(movieID int) (*Movie, error) {
	movieURL := fmt.Sprintf("%s/api%s/%d?apikey=%s", m.s.url, movieURI, movieID, m.s.apiKey)
	response, err := m.s.client.Get(movieURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = parseRadarrResponse(response)
	if err != nil {
		return nil, err
	}

	var movie Movie
	json.NewDecoder(response.Body).Decode(&movie)

	return &movie, nil
}

// ListMovies https://github.com/Radarr/Radarr/wiki/API:Movie#get
func (m *MovieService) ListMovies() (*Movies, error) {
	moviesURL := fmt.Sprintf("%s/api%s?apikey=%s", m.s.url, movieURI, m.s.apiKey)
	response, err := m.s.client.Get(moviesURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = parseRadarrResponse(response)
	if err != nil {
		return nil, err
	}

	var movies Movies
	json.NewDecoder(response.Body).Decode(&movies)

	return &movies, nil
}
