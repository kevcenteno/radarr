package main

import (
	"flag"
	"fmt"
	"runtime"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/urfave/cli/v2"
)

func Test_showVersion(t *testing.T) {
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
			out:     fmt.Sprintf("radarr %s compiled with %v on %v/%v\n", version, runtime.Version(), runtime.GOOS, runtime.GOARCH),
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
