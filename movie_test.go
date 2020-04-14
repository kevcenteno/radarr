package radarr

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func Test_newMovieService(t *testing.T) {
	s := &Service{client: http.DefaultClient, url: dummyURL}

	tests := []struct {
		name    string
		service *Service
		want    *MovieService
	}{
		{
			name:    "Constructor",
			service: s,
			want: &MovieService{
				s:      s,
				Lookup: newLookupService(s),
			},
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
	var expectedMovie Movie
	if err := json.NewDecoder(dummyMovieResponse().Body).Decode(&expectedMovie); err != nil {
		t.Fatal(err)
	}

	goodService := newMovieService(&Service{
		client: dummyHTTPClient,
		url:    dummyURL,
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
			want:    &expectedMovie,
			wantErr: false,
		},
	}

	var m *MovieService
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m = &MovieService{tt.service, nil}
			got, err := m.Get(tt.movieID)
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
	var expectedMovies Movies
	if err := json.NewDecoder(dummyMoviesResponse().Body).Decode(&expectedMovies); err != nil {
		t.Fatal(err)
	}

	goodService := newMovieService(&Service{
		client: dummyHTTPClient,
		url:    dummyURL,
	})

	tests := []struct {
		name    string
		service *Service
		want    Movies
		wantErr bool
	}{
		{
			name:    "Same response",
			service: goodService.s,
			want:    expectedMovies,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieService{tt.service, nil}
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
	var expectedMovies Movies
	if err := json.NewDecoder(dummyMoviesResponse().Body).Decode(&expectedMovies); err != nil {
		t.Fatal(err)
	}

	var expectedMovie Movie
	if err := json.NewDecoder(dummyMovieResponse().Body).Decode(&expectedMovie); err != nil {
		t.Fatal(err)
	}

	goodService := newMovieService(&Service{
		client: dummyHTTPClient,
		url:    dummyURL,
	})

	tests := []struct {
		name    string
		service *Service
		opts    []*UpcomingOptions
		want    Movies
		wantErr bool
	}{
		{
			name:    "Without filter",
			opts:    nil,
			service: goodService.s,
			want:    Movies{},
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
				return []*UpcomingOptions{{Start: &s, End: &e}}
			}(),
		},
		{
			name:    "Start filter",
			service: goodService.s,
			wantErr: false,
			want:    expectedMovies,
			opts: func() []*UpcomingOptions {
				s := time.Date(2019, time.November, 19, 23, 0, 0, 0, time.UTC)
				return []*UpcomingOptions{{Start: &s}}
			}(),
		},
		{
			name:    "End filter",
			service: goodService.s,
			want:    Movies{},
			wantErr: false,
			opts: func() []*UpcomingOptions {
				e := time.Date(2019, time.November, 20, 23, 0, 0, 0, time.UTC)
				return []*UpcomingOptions{{End: &e}}
			}(),
		},
		{
			name:    "Both filters",
			service: goodService.s,
			want:    Movies{&expectedMovie},
			wantErr: false,
			opts: func() []*UpcomingOptions {
				start := time.Date(2019, time.November, 19, 23, 0, 0, 0, time.UTC)
				end := time.Date(2019, time.November, 20, 23, 0, 0, 0, time.UTC)
				return []*UpcomingOptions{{Start: &start, End: &end}}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieService{tt.service, nil}
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

func TestMovieService_Delete(t *testing.T) {
	type args struct {
		movie *Movie
		opts  []*DeleteMovieOptions
	}

	var expectedMovie Movie
	if err := json.NewDecoder(dummyMovieResponse().Body).Decode(&expectedMovie); err != nil {
		t.Fatal(err)
	}

	goodService := newMovieService(&Service{
		client: dummyHTTPClient,
		url:    dummyURL,
	})

	tests := []struct {
		name    string
		service *Service
		args    args
		wantErr bool
	}{
		{
			name:    "Delete without option",
			service: goodService.s,
			args:    args{movie: &expectedMovie},
			wantErr: false,
		},
		{
			name:    "Delete with addExclusion option",
			args:    args{movie: &expectedMovie, opts: []*DeleteMovieOptions{{AddExclusion: true}}},
			service: goodService.s,
			wantErr: false,
		},
		{
			name:    "Delete with deleteFiles option",
			args:    args{movie: &expectedMovie, opts: []*DeleteMovieOptions{{DeleteFiles: true}}},
			service: goodService.s,
			wantErr: false,
		},
		{
			name:    "Delete with both options",
			args:    args{movie: &expectedMovie, opts: []*DeleteMovieOptions{{DeleteFiles: true, AddExclusion: true}}},
			service: goodService.s,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieService{
				s: tt.service,
			}
			if err := m.Delete(tt.args.movie, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("MovieService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMovieService_Excluded(t *testing.T) {
	var expectedMovies ExcludedMovies
	if err := json.NewDecoder(dummyExcludedMovies().Body).Decode(&expectedMovies); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		s       *Service
		want    ExcludedMovies
		wantErr bool
	}{
		{
			name: "Expected response",
			s: newMovieService(&Service{
				client: dummyHTTPClient,
				url:    dummyURL,
			}).s,
			want:    expectedMovies,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieService{
				s: tt.s,
			}
			got, err := m.Excluded()
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieService.Excluded() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieService.Excluded() = %v, want %v", got, tt.want)
			}
		})
	}
}
