package radarr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func Test_newMovieService(t *testing.T) {
	s := &Service{client: http.DefaultClient, apiKey: internal.DummyAPIKey, url: internal.DummyURL}

	tests := []struct {
		name    string
		service *Service
		want    *MovieService
	}{
		{
			name:    "Constructor",
			service: s,
			want:    &MovieService{s},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMovieService(tt.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMovieService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieService_Get(t *testing.T) {
	var expecedMovie *Movie
	err := json.Unmarshal([]byte(internal.DummyMovieResponse), &expecedMovie)
	if err != nil {
		t.Errorf("json.Unmarshal() error: %s", err.Error())
	}

	goodService := newMovieService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
		apiKey: internal.DummyAPIKey,
	})
	badService := newMovieService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
		apiKey: "foo",
	})

	tests := []struct {
		name    string
		service *Service
		movieID int
		want    *Movie
		wantErr bool
	}{
		{
			name:    "Same response",
			movieID: 217,
			service: goodService.s,
			want:    expecedMovie,
			wantErr: false,
		},
		{
			name:    "Bad API key",
			movieID: 217,
			service: badService.s,
			want:    nil,
			wantErr: true,
		},
	}

	var m *MovieService
	var got *Movie
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m = &MovieService{tt.service}
			got, err = m.Get(tt.movieID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieService_List(t *testing.T) {
	var expecedMovies *Movies
	err := json.Unmarshal([]byte(fmt.Sprintf("[%s, %s]", internal.DummyMovieResponse, internal.DummyMovieResponse)), &expecedMovies)
	if err != nil {
		t.Errorf("json.Unmarshal() error: %s", err.Error())
	}

	goodService := newMovieService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
		apiKey: internal.DummyAPIKey,
	})
	badService := newMovieService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
		apiKey: "foo",
	})

	tests := []struct {
		name    string
		service *Service
		want    *Movies
		wantErr bool
	}{
		{
			name:    "Same response",
			service: goodService.s,
			want:    expecedMovies,
			wantErr: false,
		},
		{
			name:    "Bad API key",
			service: badService.s,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieService{tt.service}
			got, err := m.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieService.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieService_Upcoming(t *testing.T) {
	var expecedMovies *Movies
	err := json.Unmarshal([]byte(fmt.Sprintf("[%s, %s]", internal.DummyMovieResponse, internal.DummyMovieResponse)), &expecedMovies)
	if err != nil {
		t.Errorf("json.Unmarshal() error: %s", err.Error())
	}

	var expecedMovie *Movie
	err = json.Unmarshal([]byte(internal.DummyMovieResponse), &expecedMovie)
	if err != nil {
		t.Errorf("json.Unmarshal() error: %s", err.Error())
	}

	goodService := newMovieService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
		apiKey: internal.DummyAPIKey,
	})
	badService := newMovieService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
		apiKey: "foo",
	})

	tests := []struct {
		name    string
		service *Service
		opts    []*UpcomingOptions
		want    *Movies
		wantErr bool
	}{
		{
			name:    "Bad API key",
			opts:    nil,
			service: badService.s,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Without filter",
			opts:    nil,
			service: goodService.s,
			want:    &Movies{},
			wantErr: false,
		},
		{
			name:    "Dates with reverse order",
			service: goodService.s,
			wantErr: true,
			want:    nil,
			opts: func() []*UpcomingOptions {
				s := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)
				e := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.Local)
				return []*UpcomingOptions{&UpcomingOptions{Start: &s, End: &e}}
			}(),
		},
		{
			name:    "Start filter",
			service: goodService.s,
			wantErr: false,
			want:    expecedMovies,
			opts: func() []*UpcomingOptions {
				s := time.Date(2019, time.November, 19, 23, 0, 0, 0, time.UTC)
				return []*UpcomingOptions{&UpcomingOptions{Start: &s}}
			}(),
		},
		{
			name:    "End filter",
			service: goodService.s,
			want:    &Movies{},
			wantErr: false,
			opts: func() []*UpcomingOptions {
				e := time.Date(2019, time.November, 20, 23, 0, 0, 0, time.UTC)
				return []*UpcomingOptions{&UpcomingOptions{End: &e}}
			}(),
		},
		{
			name:    "Both filters",
			service: goodService.s,
			want:    &Movies{*expecedMovie},
			wantErr: false,
			opts: func() []*UpcomingOptions {
				start := time.Date(2019, time.November, 19, 23, 0, 0, 0, time.UTC)
				end := time.Date(2019, time.November, 20, 23, 0, 0, 0, time.UTC)
				return []*UpcomingOptions{&UpcomingOptions{Start: &start, End: &end}}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieService{tt.service}
			got, err := m.Upcoming(tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieService.Upcoming() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieService.Upcoming() = %v, want %v", got, tt.want)
			}
		})
	}
}
