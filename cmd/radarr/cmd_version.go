package main

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

var versionCommand *cli.Command = &cli.Command{
	Name:     "version",
	Usage:    "Print version",
	HideHelp: true,
	Action:   showVersion,
}

func init() {
	app.Commands = append(app.Commands, versionCommand)
}

func showVersion(c *cli.Context) error {
	log.Debugln("Print version")
	fmt.Printf("radarr %s compiled with %v on %v/%v\n", version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	return nil
}
