package radarr

import (
	"net/http"
	"reflect"
	"testing"
)

type dummyHTTPTransport struct {
	http.RoundTripper
}

// RoundTrip mocked default HTTP client transport RoundTrip function
func (r *dummyHTTPTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func Test_newTransport(t *testing.T) {
	type args struct {
		verbose bool
		key     string
	}

	tests := []struct {
		name string
		want *transport
		args args
	}{
		{
			name: "Constructor",
			want: &transport{transport: http.DefaultTransport, apiKey: "foo", verbose: false},
			args: args{verbose: false, key: "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransport(tt.args.key, tt.args.verbose); !reflect.DeepEqual(got, tt.want) {
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
		transport: &dummyHTTPTransport{},
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
