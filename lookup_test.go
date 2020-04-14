package radarr

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func Test_newLookupService(t *testing.T) {
	service := &Service{client: http.DefaultClient, url: dummyURL}

	tests := []struct {
		name string
		s    *Service
		want *LookupService
	}{
		{
			name: "Constructor",
			s:    service,
			want: &LookupService{s: service},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newLookupService(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLookupService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupService_Plain(t *testing.T) {
	var expectedMovies Movies
	if err := json.NewDecoder(dummyMoviesResponse().Body).Decode(&expectedMovies); err != nil {
		t.Fatal(err)
	}

	service := &Service{client: dummyHTTPClient, url: dummyURL}
	tests := []struct {
		name    string
		s       *Service
		term    string
		want    Movies
		wantErr bool
	}{
		{
			name:    "Query not provided",
			s:       service,
			term:    "",
			wantErr: true,
			want:    nil,
		},
		{
			name:    "Valid request",
			s:       service,
			term:    "star wars",
			wantErr: false,
			want:    expectedMovies,
		},
		{
			name:    "Non existent movie",
			s:       service,
			term:    "this film does not exist either",
			wantErr: false,
			want:    Movies{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LookupService{
				s: tt.s,
			}
			got, err := l.Plain(tt.term)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookupService.Plain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LookupService.Plain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupService_Tmdb(t *testing.T) {
	var expectedMovie Movie
	if err := json.NewDecoder(dummyMovieResponse().Body).Decode(&expectedMovie); err != nil {
		t.Fatal(err)
	}

	service := &Service{client: dummyHTTPClient, url: dummyURL}
	tests := []struct {
		name    string
		s       *Service
		TMDBID  int
		want    *Movie
		wantErr bool
	}{
		{
			name:    "Valid ID",
			s:       service,
			want:    &expectedMovie,
			wantErr: false,
			TMDBID:  348350,
		},
		{
			name:    "Invalid ID",
			s:       service,
			want:    nil,
			wantErr: true,
			TMDBID:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LookupService{
				s: tt.s,
			}
			got, err := l.Tmdb(tt.TMDBID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookupService.Tmdb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LookupService.Tmdb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupService_Imdb(t *testing.T) {
	var expectedMovie Movie
	if err := json.NewDecoder(dummyMovieResponse().Body).Decode(&expectedMovie); err != nil {
		t.Fatal(err)
	}

	service := &Service{client: dummyHTTPClient, url: dummyURL}
	tests := []struct {
		name    string
		s       *Service
		IMDBID  string
		want    *Movie
		wantErr bool
	}{
		{
			name:    "Valid ID",
			s:       service,
			want:    &expectedMovie,
			wantErr: false,
			IMDBID:  "tt3778644",
		},
		{
			name:    "Invalid ID",
			s:       service,
			want:    nil,
			wantErr: true,
			IMDBID:  "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LookupService{
				s: tt.s,
			}
			got, err := l.Imdb(tt.IMDBID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookupService.Imdb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LookupService.Imdb() = %v, want %v", got, tt.want)
			}
		})
	}
}
