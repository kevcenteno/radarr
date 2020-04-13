package main

import (
	"flag"
	"fmt"
	"runtime"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/urfave/cli/v2"
)

func Test_version_String(t *testing.T) {
	type fields struct {
		GOOS    string
		GOARCH  string
		Version string
		Runtime string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Format is well",
			fields: fields{GOARCH: "foo", GOOS: "bar", Runtime: "1", Version: "0"},
			want:   "radarr 0 compiled with 1 on bar/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &version{
				GOOS:    tt.fields.GOOS,
				GOARCH:  tt.fields.GOARCH,
				Version: tt.fields.Version,
				Runtime: tt.fields.Runtime,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_showVersion(t *testing.T) {
	// TODO: Need help to test the --json flag
	var c *cli.Context = cli.NewContext(cli.NewApp(), flag.CommandLine, nil)
	tests := []struct {
		name    string
		c       *cli.Context
		wantErr bool
		out     string
	}{
		{
			name:    "Plain",
			wantErr: false,
			c:       c,
			out:     fmt.Sprintf("radarr %s compiled with %v on %v/%v\n", Version, runtime.Version(), runtime.GOOS, runtime.GOARCH),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := capturer.CaptureStdout(func() {
				if err := showVersion(tt.c); (err != nil) != tt.wantErr {
					t.Errorf("showVersion() error = %v, wantErr %v", err, tt.wantErr)
				}
			})

			if out != tt.out {
				t.Errorf("showVersion() error = %v, want %v", out, tt.out)
			}
		})
	}
}
