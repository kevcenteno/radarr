package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/SkYNewZ/radarr"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var radarrClient *radarr.Service

var t *tablewriter.Table = tablewriter.NewWriter(os.Stdout)

var version string = "development"
var verbose bool = false
var log *logrus.Logger = logrus.New()

var app *cli.App = &cli.App{
	Name:                 "Radarr CLI",
	Usage:                "Perform actions on your Radarr instance",
	Version:              version,
	Compiled:             time.Now(),
	EnableBashCompletion: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "url",
			EnvVars:  []string{"RADARR_URL"},
			Required: true,
			Usage:    "Radarr instance URL",
		},
		&cli.StringFlag{
			Name:     "apiKey",
			EnvVars:  []string{"RADARR_API_KEY"},
			Required: true,
			Usage:    "Radarr API key",
		},
		&cli.BoolFlag{
			Name:  "json",
			Usage: "Print output as JSON instead of table",
		},
		&cli.BoolFlag{
			Name:        "verbose",
			Aliases:     []string{"v"},
			Usage:       "Print debug infos",
			Destination: &verbose,
		},
	},
	HideVersion: true,
	Authors:     []*cli.Author{{Email: "quentin@lemairepro.fr", Name: "SkYNewZ"}},
	Before: func(c *cli.Context) error {
		// Instantiate Radarr client
		s, err := radarr.New(c.String("url"), c.String("apiKey"), nil, &radarr.ClientOptions{
			Verbose: verbose,
		})
		if err != nil {
			return err
		}

		radarrClient = s

		// Set debug level
		if verbose {
			log.SetLevel(logrus.DebugLevel)
		}
		return nil
	},
}

func init() {
	t.SetAlignment(tablewriter.ALIGN_LEFT)
	t.SetAutoWrapText(false)
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
