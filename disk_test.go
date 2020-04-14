package radarr

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func Test_newDiskspaceService(t *testing.T) {
	type args struct {
		s *Service
	}

	s := &Service{client: http.DefaultClient, url: dummyURL}
	tests := []struct {
		name string
		args args
		want *DiskspaceService
	}{
		{
			name: "Constructor",
			args: args{s},
			want: &DiskspaceService{s},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDiskspaceService(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDiskspaceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiskspaceService_Get(t *testing.T) {
	var exepectedResponse *Diskspaces
	if err := json.NewDecoder(dummyDiskspaceResponse().Body).Decode(&exepectedResponse); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		s       *Service
		want    *Diskspaces
		wantErr bool
	}{
		{
			name:    "Diskspace lengh should be 2",
			s:       &Service{client: dummyHTTPClient, url: dummyURL},
			wantErr: false,
			want:    exepectedResponse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DiskspaceService{
				s: tt.s,
			}
			got, err := s.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("DiskspaceService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiskspaceService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
