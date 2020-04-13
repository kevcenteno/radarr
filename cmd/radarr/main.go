package main

import (
	"log"
	"os"
	"sort"
	"time"

	"github.com/SkYNewZ/radarr"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

var radarURL string
var radarrAPIKey string
var radarrClient *radarr.Service

var t *tablewriter.Table = tablewriter.NewWriter(os.Stdout)

// Version program version
var Version string = "development"

var app *cli.App = &cli.App{
	Name:                 "Radarr CLI",
	Usage:                "Perform actions on your Radarr instance",
	Version:              Version,
	Compiled:             time.Now(),
	EnableBashCompletion: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "url",
			EnvVars:     []string{"RADARR_URL"},
			Required:    true,
			Usage:       "Radarr instance URL",
			Destination: &radarURL,
		},
		&cli.StringFlag{
			Name:        "apiKey",
			EnvVars:     []string{"RADARR_API_KEY"},
			Required:    true,
			Usage:       "Radarr API key",
			Destination: &radarrAPIKey,
		},
	},
	Authors: []*cli.Author{{Email: "quentin@lemairepro.fr", Name: "SkYNewZ"}},
	Before: func(c *cli.Context) error {
		s, err := radarr.New(radarURL, radarrAPIKey, nil)
		if err != nil {
			return err
		}

		radarrClient = s
		return nil
	},
}

func init() {
	log.SetFlags(0)
	t.SetAlignment(tablewriter.ALIGN_LEFT)
	t.SetAutoWrapText(false)
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
