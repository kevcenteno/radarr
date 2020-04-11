package radarr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// History return your Radarr history
type History struct {
	Page          int      `json:"page"`
	PageSize      int      `json:"pageSize"`
	SortKey       string   `json:"sortKey"`
	SortDirection string   `json:"sortDirection"`
	TotalRecords  int      `json:"totalRecords"`
	Records       []Record `json:"records"`
}

// Revision movie revisions
type Revision struct {
	Version  int  `json:"version"`
	Real     int  `json:"real"`
	IsRepack bool `json:"isRepack"`
}

// Images movie images
type Images struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url"`
}

// Ratings your movies ratings
type Ratings struct {
	Votes int     `json:"votes"`
	Value float64 `json:"value"`
}

// Data contains arbitrary data
type Data struct {
	Reason          string    `json:"reason,omitempty"`
	DroppedPath     string    `json:"droppedPath,omitempty"`
	ImportedPath    string    `json:"importedPath,omitempty"`
	Indexer         string    `json:"indexer,omitempty"`
	NzbInfoURL      string    `json:"nzbInfoUrl,omitempty"`
	ReleaseGroup    string    `json:"releaseGroup,omitempty"`
	Age             string    `json:"age,omitempty"`
	AgeHours        string    `json:"ageHours,omitempty"`
	AgeMinutes      string    `json:"ageMinutes,omitempty"`
	PublishedDate   time.Time `json:"publishedDate,omitempty"`
	DownloadClient  string    `json:"downloadClient,omitempty"`
	Size            string    `json:"size,omitempty"`
	DownloadURL     string    `json:"downloadUrl,omitempty"`
	GUID            string    `json:"guid,omitempty"`
	TvdbID          string    `json:"tvdbId,omitempty"`
	TvRageID        string    `json:"tvRageId,omitempty"`
	Protocol        string    `json:"protocol,omitempty"`
	IndexerFlags    string    `json:"indexerFlags,omitempty"`
	IndexerID       string    `json:"indexerId,omitempty"`
	TorrentInfoHash string    `json:"torrentInfoHash,omitempty"`
}

// Record history item
type Record struct {
	MovieID             int       `json:"movieId"`
	SourceTitle         string    `json:"sourceTitle"`
	Quality             Quality   `json:"quality"`
	QualityCutoffNotMet bool      `json:"qualityCutoffNotMet"`
	Date                time.Time `json:"date"`
	EventType           string    `json:"eventType"`
	Movie               Movie     `json:"movie"`
	ID                  int       `json:"id"`
	DownloadID          string    `json:"downloadId,omitempty"`
	Data                Data      `json:"data,omitempty"`
}

// Records is a set of Record
type Records = []Record

// HistoryService perform actions on your Radarr history
type HistoryService struct {
	s *Service
}

func newHistoryService(s *Service) *HistoryService {
	return &HistoryService{s: s}
}

// Get return all history
func (s *HistoryService) Get() (*Records, error) {

	// First call
	var page int = 1
	history, err := s.paginate(page)
	if err != nil {
		return nil, err
	}

	for {
		if len(history.Records) == history.TotalRecords {
			break
		}

		page++
		s, err := s.paginate(page)
		if err != nil {
			break
		}

		history.Records = append(history.Records, s.Records...)
	}

	return &history.Records, nil
}

func (s *HistoryService) paginate(page int) (*History, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api%s", s.s.url, historyURI), nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	q.Add("page", strconv.Itoa(page))
	q.Add("pageSize", "50")

	request.URL.RawQuery = q.Encode()
	response, err := s.s.client.Do(request)
	if err != nil {
		return nil, err
	}
	err = parseRadarrResponse(response)
	if err != nil {
		return nil, err
	}

	var history History
	err = json.NewDecoder(response.Body).Decode(&history)
	if err != nil {
		return nil, err
	}
	return &history, nil
}
