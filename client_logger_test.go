package radarr

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
	"github.com/kami-zh/go-capturer"
)

type test struct {
	title    string
	expected interface{}
	got      interface{}
}

func TestLogRequestMessage(t *testing.T) {
	m := fmt.Sprintf(logReqMsg, "foo")
	cases := []test{
		test{
			title:    "Message should contain 'foo'",
			expected: true,
			got:      strings.Contains(m, "foo"),
		},
		test{
			title:    "Message should contain 'API Request Details:'",
			expected: true,
			got:      strings.Contains(m, "API Request Details:"),
		},
		test{
			title:    "Message should contain '[ REQUEST ]'",
			expected: true,
			got:      strings.Contains(m, "[ REQUEST ]"),
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			if c.expected != c.got {
				t.Errorf("Got '%v' want '%v'", c.got, c.expected)
			}
		})
	}
}

func TestLogResponseMessage(t *testing.T) {
	m := fmt.Sprintf(logRespMsg, "foo")
	cases := []test{
		test{
			title:    "Message should contain 'foo'",
			expected: true,
			got:      strings.Contains(m, "foo"),
		},
		test{
			title:    "Message should contain 'API Response Details:'",
			expected: true,
			got:      strings.Contains(m, "API Response Details:"),
		},
		test{
			title:    "Message should contain '[ RESPONSE ]'",
			expected: true,
			got:      strings.Contains(m, "[ RESPONSE ]"),
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			if c.expected != c.got {
				t.Errorf("Got '%v' want '%v'", c.got, c.expected)
			}
		})
	}
}

func TestTransport_RoundTrip(t *testing.T) {
	// Request send to internal.MockedTransports
	var exepectedRequest = http.Request{Header: http.Header{}, URL: internal.ParseDummyURL}

	var expectedDebugLog = []string{
		"API Request Details:",
		"GET / HTTP/1.1",
		fmt.Sprintf("Host: %s", internal.ParseDummyURL.Host),
		"User-Agent: SkYNewZ-Go-http-client/1.1",
		"Content-Type: application/json; charset=utf-8",
		"Accept-Encoding: gzip",
		"API Response Details:",
		"HTTP/0.0 200 OK",
		`{"foo": "bar"}`,
	}

	var mockedTransports = internal.NewMockedTransports()

	// Create a new Transport to override default one
	trans := transport{
		transport: mockedTransports.MockedTransport1,
	}

	// Execute function without DEBUG
	var err error
	var response *http.Response
	out := capturer.CaptureStdout(func() {
		_ = os.Unsetenv(envLog)
		response, err = trans.RoundTrip(&exepectedRequest)
	})

	cases := []test{
		test{
			title:    "Error should be nil",
			expected: nil,
			got:      err,
		},
		test{
			title:    "Request header should contain correct User-Agent",
			expected: "SkYNewZ-Go-http-client/1.1",
			got:      exepectedRequest.Header.Get("User-Agent"),
		},
		test{
			title:    "Request header should contain correct Content-Type",
			expected: "application/json; charset=utf-8",
			got:      exepectedRequest.Header.Get("Content-Type"),
		},
		test{
			title:    fmt.Sprintf("%s is NOT set: nothing should be print", envLog),
			expected: "",
			got:      out,
		},
		test{
			title:    "Response should be the same",
			expected: mockedTransports.MockedTransport1.MockedResponse,
			got:      response,
		},
	}

	// Execute function with DEBUG
	out = capturer.CaptureStdout(func() {
		_ = os.Setenv(envLog, "DEBUG")
		_, err = trans.RoundTrip(&exepectedRequest)
		_ = os.Unsetenv(envLog)
	})

	for _, s := range expectedDebugLog {
		cases = append(cases, test{
			title:    fmt.Sprintf("%s is set: debug request/response should be equal", envLog),
			expected: true,
			got:      strings.Contains(out, s),
		})
	}

	// Test returned error
	trans = transport{
		transport: mockedTransports.MockedTransport2,
	}
	_, err = trans.RoundTrip(&exepectedRequest)
	cases = append(cases, test{
		title:    "Error should be not nil",
		expected: false,
		got:      err == nil,
	})

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			if c.expected != c.got {
				t.Errorf("Got '%v' want '%v'", c.got, c.expected)
			}
		})
	}
}

func Test_newTransport(t *testing.T) {
	tests := []struct {
		name string
		want *transport
	}{
		{
			name: "Constructor",
			want: &transport{transport: http.DefaultTransport},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransport(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDebugLevel(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "DEBUG",
			want: true,
		},
		{
			name: "Foo",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(envLog, tt.name)
			if got := isDebugLevel(); got != tt.want {
				t.Errorf("isDebugLevel() = %v, want %v", got, tt.want)
			}
		})
	}
	os.Unsetenv(envLog)
}
