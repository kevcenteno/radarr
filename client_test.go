package radarr

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func TestNew(t *testing.T) {
	type args struct {
		radarrURL string
		apiKey    string
		client    HTTPClientInterface
	}

	var serviceWithCustomHTTPClient *Service = &Service{url: internal.DummyURL, apiKey: internal.DummyAPIKey, client: internal.DummyHTTPClient}
	serviceWithCustomHTTPClient.Movies = newMovieService(serviceWithCustomHTTPClient)
	serviceWithCustomHTTPClient.Diskspace = newDiskspaceService(serviceWithCustomHTTPClient)
	serviceWithCustomHTTPClient.SystemStatus = newSystemStatusService(serviceWithCustomHTTPClient)
	serviceWithCustomHTTPClient.Command = newCommandService(serviceWithCustomHTTPClient)

	client := http.Client{}
	client.Timeout = time.Second * 10
	client.Transport = newTransport()
	var serviceWithDefaultHTTPClient *Service = &Service{url: internal.DummyURL, apiKey: internal.DummyAPIKey, client: &client}
	serviceWithDefaultHTTPClient.Movies = newMovieService(serviceWithDefaultHTTPClient)
	serviceWithDefaultHTTPClient.Diskspace = newDiskspaceService(serviceWithDefaultHTTPClient)
	serviceWithDefaultHTTPClient.SystemStatus = newSystemStatusService(serviceWithDefaultHTTPClient)
	serviceWithDefaultHTTPClient.Command = newCommandService(serviceWithDefaultHTTPClient)

	tests := []struct {
		name    string
		args    args
		want    *Service
		wantErr bool
	}{
		struct {
			name    string
			args    args
			want    *Service
			wantErr bool
		}{
			name:    "Error because of bad URL",
			args:    args{apiKey: internal.DummyAPIKey, radarrURL: "bad-url", client: internal.DummyHTTPClient},
			wantErr: true,
		},
		struct {
			name    string
			args    args
			want    *Service
			wantErr bool
		}{
			name:    "Good service",
			args:    args{radarrURL: internal.DummyURL, apiKey: internal.DummyAPIKey, client: internal.DummyHTTPClient},
			wantErr: false,
			want:    serviceWithCustomHTTPClient,
		},
		struct {
			name    string
			args    args
			want    *Service
			wantErr bool
		}{
			name:    "Default HTTP Client",
			args:    args{radarrURL: internal.DummyURL, apiKey: internal.DummyAPIKey, client: nil},
			wantErr: false,
			want:    serviceWithDefaultHTTPClient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.radarrURL, tt.args.apiKey, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
