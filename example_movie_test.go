package radarr_test

import (
	"fmt"
	"log"
	"time"

	"github.com/SkYNewZ/radarr"
)

// Return upcoming movies by specifying period
func ExampleMovieService_Upcoming_basic() {
	client, err := radarr.New("https://my.radarr-instance.fr", "radarr-api-key", nil)
	if err != nil {
		log.Fatalln(err)
	}

	movies, err := client.Movies.Upcoming()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v", movies)
}

func ExampleMovieService_Upcoming_advanced() {
	client, err := radarr.New("https://my.radarr-instance.fr", "radarr-api-key", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Search upcoming movies between April 5th and April 10th")
	start := time.Date(2020, time.April, 5, 0, 0, 0, 0, time.Local)
	end := time.Date(2020, time.April, 10, 0, 0, 0, 0, time.Local)
	movies, err := client.Movies.Upcoming(&radarr.UpcomingOptions{
		Start: &start,
		End:   &end,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v", movies)
}
