package radarr

import (
	"encoding/json"
	"fmt"
	"time"
)

// SystemStatus Radarr system status response
type SystemStatus struct {
	Version           string    `json:"version"`
	BuildTime         time.Time `json:"buildTime"`
	IsDebug           bool      `json:"isDebug"`
	IsProduction      bool      `json:"isProduction"`
	IsAdmin           bool      `json:"isAdmin"`
	IsUserInteractive bool      `json:"isUserInteractive"`
	StartupPath       string    `json:"startupPath"`
	AppData           string    `json:"appData"`
	OsName            string    `json:"osName"`
	OsVersion         string    `json:"osVersion"`
	IsNetCore         bool      `json:"isNetCore"`
	IsMono            bool      `json:"isMono"`
	IsLinux           bool      `json:"isLinux"`
	IsOsx             bool      `json:"isOsx"`
	IsWindows         bool      `json:"isWindows"`
	Branch            string    `json:"branch"`
	Authentication    string    `json:"authentication"`
	SqliteVersion     string    `json:"sqliteVersion"`
	MigrationVersion  int       `json:"migrationVersion"`
	URLBase           string    `json:"urlBase"`
	RuntimeVersion    string    `json:"runtimeVersion"`
	RuntimeName       string    `json:"runtimeName"`
}

// SystemStatusService contains Radarr system operations
type SystemStatusService struct {
	s *Service
}

func newSystemStatusService(s *Service) *SystemStatusService {
	return &SystemStatusService{s}
}

// Get https://github.com/Radarr/Radarr/wiki/API:System-Status#get
func (s *SystemStatusService) Get() (*SystemStatus, error) {
	statusURL := fmt.Sprintf("%s/api%s", s.s.url, statusURI)
	resp, err := s.s.client.Get(statusURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = parseRadarrResponse(resp)
	if err != nil {
		return nil, err
	}

	var status SystemStatus
	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}
