// Radarr CLI. Perform actions on your Raddarr instance using CLI
//
// Installation
//
// Simply `go get` this package
// 	go get github.com/SkYNewZ/radarr/cmd/radarr
//
// List movies
//
// List all your movies
//  radarr --url "https://my.radarr-instance.fr" --apiKey "radarr-api-key" movies ls
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/SkYNewZ/radarr"
	"github.com/urfave/cli/v2"
)

var radarURL string
var radarrAPIKey string

// Version program version
var Version string = "development"

func main() {
	log.SetFlags(0)

	app := &cli.App{
		Name:     "Radarr CLI",
		Usage:    "Perform actions on your Radarr instance using CLI",
		Version:  Version,
		Compiled: time.Now(),
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
		Authors: []*cli.Author{&cli.Author{Email: "quentin@lemairepro.fr", Name: "SkYNewZ"}},
		Commands: []*cli.Command{
			&cli.Command{
				Name:  "movies",
				Usage: "Perform actions on movies",
				Subcommands: []*cli.Command{
					&cli.Command{
						Name:    "list",
						Usage:   "List all movies in your collection",
						Aliases: []string{"ls"},
						Action:  listMovies,
					},
					&cli.Command{
						Name:   "get",
						Usage:  "Search a movie bu ID",
						Action: getMovie,
					},
					&cli.Command{
						Name:   "upcoming",
						Usage:  "List upcoming movies",
						Action: upcoming,
						Flags: []cli.Flag{
							&cli.TimestampFlag{
								Name:        "start",
								Required:    false,
								Usage:       "Specify a start date",
								Layout:      "2006-01-02T15:04:05Z",
								DefaultText: "",
							},
							&cli.TimestampFlag{
								Name:        "end",
								Required:    false,
								Usage:       "Specify a end date",
								Layout:      "2006-01-02T15:04:05Z",
								DefaultText: "",
							},
						},
					},
				},
			},
			&cli.Command{
				Name:   "status",
				Usage:  "Get your Radarr instance status",
				Action: getStatus,
			},
		},
	}

	app.EnableBashCompletion = true
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initRadarrClient() (*radarr.Service, error) {
	s, err := radarr.New(radarURL, radarrAPIKey, nil)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func listMovies(*cli.Context) error {
	client, err := initRadarrClient()
	if err != nil {
		return err
	}

	movies, err := client.Movies.List()
	if err != nil {
		return err
	}

	r, err := json.Marshal(movies)
	if err != nil {
		return err
	}

	fmt.Println(string(r))
	return nil
}

func getMovie(c *cli.Context) error {
	client, err := initRadarrClient()
	if err != nil {
		return err
	}

	movieID := c.Args().First()
	m, err := strconv.Atoi(movieID)
	if err != nil {
		return err
	}

	movie, err := client.Movies.Get(m)
	if err != nil {
		return err
	}

	r, err := json.Marshal(movie)
	if err != nil {
		return err
	}

	fmt.Println(string(r))
	return nil
}

func getStatus(*cli.Context) error {
	client, err := initRadarrClient()
	if err != nil {
		return err
	}

	status, err := client.SystemStatus.Get()
	if err != nil {
		return err
	}

	r, err := json.Marshal(status)
	if err != nil {
		return err
	}

	fmt.Println(string(r))
	return nil
}

func upcoming(c *cli.Context) error {
	client, err := initRadarrClient()
	if err != nil {
		return err
	}

	start := c.Value("start").(cli.Timestamp)
	end := c.Value("end").(cli.Timestamp)

	opts := &radarr.UpcomingOptions{}
	if start.Value() != nil {
		opts.Start = start.Value()
	}
	if end.Value() != nil {
		opts.End = end.Value()
	}

	movies, err := client.Movies.Upcoming(opts)
	if err != nil {
		return err
	}

	r, err := json.Marshal(movies)
	if err != nil {
		return err
	}

	fmt.Println(string(r))
	return nil
}
