# Radarr Go client

[![](https://github.com/SkYNewZ/radarr/workflows/CD/badge.svg)](https://github.com/SkYNewZ/radarr/actions)
[![](https://gocover.io/_badge/github.com/skynewz/radarr)](https://gocover.io/github.com/SkYNewZ/radarr)
[![Go Report Card](https://goreportcard.com/badge/github.com/SkYNewZ/radarr)](https://goreportcard.com/report/github.com/SkYNewZ/radarr)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/skynewz/radarr)
[![Godoc](https://godoc.org/github.com/SkYNewZ/radarr?status.svg)](https://godoc.org/github.com/SkYNewZ/radarr)
[![Docker Pulls](https://img.shields.io/docker/pulls/skynewz/radarr)](https://hub.docker.com/r/skynewz/radarr)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/skynewz/radarr)](https://github.com/SkYNewZ/radarr/releases)

This is a Go package that lets you to interact with your Radarr instance.
Developed for [Radarr API v2](https://github.com/Radarr/Radarr/wiki/API).

Radarr API v3 is under construction. See [here](https://github.com/orgs/Radarr/projects/2) and [here](https://www.reddit.com/r/radarr/comments/ejjiw2/api_v3/).

You can use it as CLI. See [related section](cmd/radarr)

## Supports

Here are the currently supported endpoints:

- [x] Calendar
- [ ] Command
- [x] Diskspace
- [x] History
- [ ] Movie
  - [x] Returns all Movies in your collection
  - [x] Returns the movie with the matching ID or 404 if no matching movie is found
  - [ ] Adds a new movie to your collection
  - [ ] Update an existing movie
  - [x] Delete the movie with the given ID
- [x] Movie Lookup
- [ ] Queue
- [x] List Exclusions
- [x] System-Status

## Getting started

```go
package main

import (
	"fmt"
	"log"

	"github.com/SkYNewZ/radarr"
)

// Instantiate a standard client
func main() {
	client, err := radarr.New("https://my.radarr-instance.fr", "radarr-api-key", nil)
	if err != nil {
		log.Fatalln(err)
	}

	movie, err := client.Movies.Get(217)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", movie.Title)

	// Output:
	// Frozen II
}
```
