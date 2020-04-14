package radarr

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func Test_newSystemStatusService(t *testing.T) {
	type args struct {
		s *Service
	}
	s := &Service{client: http.DefaultClient, url: dummyURL}
	tests := []struct {
		name string
		args args
		want *SystemStatusService
	}{
		{
			name: "Constructor",
			args: args{s},
			want: &SystemStatusService{s},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSystemStatusService(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSystemStatusService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemStatusService_Get(t *testing.T) {
	var expectedStatus *SystemStatus
	if err := json.NewDecoder(dummySystemStatusResponse().Body).Decode(&expectedStatus); err != nil {
		t.Fatal(err)
	}

	goodService := newSystemStatusService(&Service{
		client: dummyHTTPClient,
		url:    dummyURL,
	})

	tests := []struct {
		name    string
		service *Service
		want    *SystemStatus
		wantErr bool
	}{
		{
			name:    "Same response",
			service: goodService.s,
			want:    expectedStatus,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SystemStatusService{tt.service}
			got, err := s.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("SystemStatusService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemStatusService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
