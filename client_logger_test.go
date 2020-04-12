package radarr

import (
	"net/http"
	"reflect"
	"testing"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func Test_newTransport(t *testing.T) {
	tests := []struct {
		name string
		want *transport
		key  string
	}{
		{
			name: "Constructor",
			want: &transport{transport: http.DefaultTransport, apiKey: "foo"},
			key:  "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransport(tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transport_RoundTrip(t *testing.T) {
	req := &http.Request{Header: http.Header{}}
	tests := []struct {
		name string
		req  *http.Request
		want string
	}{
		{
			name: "X-Api-Key",
			req:  req,
			want: "foo",
		},
		{
			name: "Content-Type",
			req:  req,
			want: "application/json; charset=utf-8",
		},
		{
			name: "User-Agent",
			req:  req,
			want: "SkYNewZ-Go-http-client/1.1",
		},
	}

	// Fake transport to avoid the real HTTP request
	trans := transport{
		transport: &internal.DummyHTTPTransport{},
		apiKey:    "foo",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = trans.RoundTrip(tt.req)
			if req.Header.Get(tt.name) != tt.want {
				t.Errorf("transport.RoundTrip() = %s, want %v", req.Header.Get(tt.name), tt.want)
			}
		})
	}
}
