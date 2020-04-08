package radarr

import (
	"encoding/json"
	"fmt"
)

// Diskspace disk space Radarr response
type Diskspace struct {
	Path       string `json:"path"`
	Label      string `json:"label"`
	FreeSpace  int64  `json:"freeSpace"`
	TotalSpace int64  `json:"totalSpace"`
}

// Diskspaces describe disk space info on your Radarr instance
type Diskspaces []Diskspace

// DiskspaceService contains Radarr diskspace operations
type DiskspaceService struct {
	s *Service
}

func newDiskspaceService(s *Service) *DiskspaceService {
	return &DiskspaceService{
		s: s,
	}
}

// Get return Radarr disk space info
func (s *DiskspaceService) Get() (*Diskspaces, error) {
	diskspaceURL := fmt.Sprintf("%s/api%s?apikey=%s", s.s.url, diskspaceURI, s.s.apiKey)
	response, err := s.s.client.Get(diskspaceURL)
	if err != nil {
		return nil, err
	}

	err = parseRadarrResponse(response)
	if err != nil {
		return nil, err
	}

	var diskspaces Diskspaces
	err = json.NewDecoder(response.Body).Decode(&diskspaces)
	if err != nil {
		return nil, err
	}

	_ = response.Body.Close()
	return &diskspaces, nil
}
