# Radarr Go client

![](https://github.com/SkYNewZ/radarr/workflows/Release%20package/badge.svg)

This is a Go package that lets you to interact with your Radarr instance.
Developed for [Radarr API v2](https://github.com/Radarr/Radarr/wiki/API).

Radarr API v3 is under construction. See [here](https://github.com/orgs/Radarr/projects/2) and [here](https://www.reddit.com/r/radarr/comments/ejjiw2/api_v3/).

You can use it as CLI. See [related section](cmd/radarr)

## Supports

Here are the currently supported endpoints:

- [x] Calendar
- [ ] Command
- [ ] Diskspace
- [ ] History
- [ ] Movie
  - [x] Returns all Movies in your collection
  - [x] Returns the movie with the matching ID or 404 if no matching movie is found
  - [ ] Adds a new movie to your collection
  - [ ] Update an existing movie
  - [ ] Delete the movie with the given ID
- [ ] Movie Lookup
- [ ] Queue
- [x] System-Status

## Getting started

```go
package main

import (
	"fmt"
	"log"

	"github.com/SkYNewZ/radarr"
)

// Instanciate a standard client
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

## Debug

You can set `LOG_LEVEL=DEBUG` to print all requests/responses
