package radarr

import (
	"net/http"
	"reflect"
	"testing"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func Test_newDiskspaceService(t *testing.T) {
	type args struct {
		s *Service
	}

	s := &Service{client: http.DefaultClient, apiKey: internal.DummyAPIKey, url: internal.DummyURL}
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
	type fields struct {
		s *Service
	}

	var classicService *Service = &Service{client: internal.DummyHTTPClient, url: internal.DummyURL, apiKey: internal.DummyAPIKey}
	var badAPIKeyService *Service = &Service{client: internal.DummyHTTPClient, url: internal.DummyURL, apiKey: "bad"}

	var exepectedResponse *Diskspaces = &Diskspaces{
		Diskspace{
			Label:      "/",
			Path:       "/",
			FreeSpace:  11059175424,
			TotalSpace: 15614754816,
		},
		Diskspace{
			Label:      "/home",
			Path:       "/home",
			FreeSpace:  88775757824,
			TotalSpace: 98325770240,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		want    *Diskspaces
		wantErr bool
	}{
		{
			name:    "Diskspace lengh should be 2",
			fields:  fields{classicService},
			wantErr: false,
			want:    exepectedResponse,
		},
		{
			name:    "Bad API Key",
			wantErr: true,
			fields:  fields{badAPIKeyService},
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DiskspaceService{
				s: tt.fields.s,
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
