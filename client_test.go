package radarr

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		radarrURL string
		apiKey    string
		client    HTTPClientInterface
	}

	var serviceWithCustomHTTPClient *Service = &Service{url: dummyURL, client: dummyHTTPClient}
	serviceWithCustomHTTPClient.Movies = newMovieService(serviceWithCustomHTTPClient)
	serviceWithCustomHTTPClient.Diskspace = newDiskspaceService(serviceWithCustomHTTPClient)
	serviceWithCustomHTTPClient.SystemStatus = newSystemStatusService(serviceWithCustomHTTPClient)
	serviceWithCustomHTTPClient.Command = newCommandService(serviceWithCustomHTTPClient)
	serviceWithCustomHTTPClient.History = newHistoryService(serviceWithCustomHTTPClient)

	client := http.Client{}
	client.Timeout = time.Second * 10
	client.Transport = newTransport(dummyAPIKey)
	var serviceWithDefaultHTTPClient *Service = &Service{url: dummyURL, client: &client}
	serviceWithDefaultHTTPClient.Movies = newMovieService(serviceWithDefaultHTTPClient)
	serviceWithDefaultHTTPClient.Diskspace = newDiskspaceService(serviceWithDefaultHTTPClient)
	serviceWithDefaultHTTPClient.SystemStatus = newSystemStatusService(serviceWithDefaultHTTPClient)
	serviceWithDefaultHTTPClient.Command = newCommandService(serviceWithDefaultHTTPClient)
	serviceWithDefaultHTTPClient.History = newHistoryService(serviceWithDefaultHTTPClient)

	tests := []struct {
		name    string
		args    args
		want    *Service
		wantErr bool
	}{
		{
			name:    "Error because of bad URL",
			args:    args{apiKey: dummyAPIKey, radarrURL: "bad-url", client: dummyHTTPClient},
			wantErr: true,
		},
		{
			name:    "Error because of non-provided API key",
			args:    args{radarrURL: dummyURL, client: dummyHTTPClient},
			wantErr: true,
		},
		{
			name:    "Error because of non-provided API key",
			args:    args{apiKey: "", radarrURL: dummyURL, client: dummyHTTPClient},
			wantErr: true,
		},
		{
			name:    "Good service",
			args:    args{radarrURL: dummyURL, apiKey: dummyAPIKey, client: dummyHTTPClient},
			wantErr: false,
			want:    serviceWithCustomHTTPClient,
		},
		{
			name:    "Default HTTP Client",
			args:    args{radarrURL: dummyURL, apiKey: dummyAPIKey, client: nil},
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
