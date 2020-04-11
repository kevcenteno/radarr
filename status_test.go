package radarr

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func Test_newSystemStatusService(t *testing.T) {
	type args struct {
		s *Service
	}
	s := &Service{client: http.DefaultClient, url: internal.DummyURL}
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
	var expectedStatus *SystemStatus = &SystemStatus{
		Version:           "3.0.0.2741",
		BuildTime:         time.Date(2020, time.March, 23, 16, 23, 16, 00, time.UTC),
		IsDebug:           false,
		IsProduction:      true,
		IsAdmin:           false,
		IsUserInteractive: false,
		StartupPath:       "/opt/radarr",
		AppData:           "/config",
		OsName:            "ubuntu",
		OsVersion:         "20.04",
		IsNetCore:         true,
		IsMono:            false,
		IsLinux:           true,
		IsOsx:             false,
		IsWindows:         false,
		Branch:            "develop",
		Authentication:    "forms",
		SqliteVersion:     "3.31.1",
		MigrationVersion:  169,
		URLBase:           "",
		RuntimeVersion:    "3.1.2",
		RuntimeName:       "netCore",
	}

	goodService := newSystemStatusService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
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
