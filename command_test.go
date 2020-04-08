package radarr

import (
	"reflect"
	"testing"
)

func Test_newCommandService(t *testing.T) {
	type args struct {
		s *Service
	}
	tests := []struct {
		name string
		args args
		want *CommandService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newCommandService(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCommandService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_Status(t *testing.T) {
	type fields struct {
		s *Service
	}
	type args struct {
		commandID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.Status(tt.args.commandID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_StatusAll(t *testing.T) {
	type fields struct {
		s *Service
	}
	tests := []struct {
		name   string
		fields fields
		want   *Tasks
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.StatusAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.StatusAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_RefreshMovie(t *testing.T) {
	type fields struct {
		s *Service
	}
	type args struct {
		movieID []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.RefreshMovie(tt.args.movieID...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.RefreshMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_RescanMovie(t *testing.T) {
	type fields struct {
		s *Service
	}
	type args struct {
		movieID []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.RescanMovie(tt.args.movieID...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.RescanMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_MoviesSearch(t *testing.T) {
	type fields struct {
		s *Service
	}
	type args struct {
		movieIDs [][]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.MoviesSearch(tt.args.movieIDs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.MoviesSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_DownloadedMoviesScan(t *testing.T) {
	type fields struct {
		s *Service
	}
	type args struct {
		opts []*DownloadedMoviesScanOptions
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.DownloadedMoviesScan(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.DownloadedMoviesScan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_RssSync(t *testing.T) {
	type fields struct {
		s *Service
	}
	tests := []struct {
		name   string
		fields fields
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.RssSync(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.RssSync() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_RenameFiles(t *testing.T) {
	type fields struct {
		s *Service
	}
	type args struct {
		files [][]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.RenameFiles(tt.args.files...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.RenameFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_RenameMovie(t *testing.T) {
	type fields struct {
		s *Service
	}
	type args struct {
		movieIDs [][]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.RenameMovie(tt.args.movieIDs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.RenameMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_CutOffUnmetMoviesSearch(t *testing.T) {
	type fields struct {
		s *Service
	}
	type args struct {
		filter *Filter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.CutOffUnmetMoviesSearch(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.CutOffUnmetMoviesSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_NetImportSync(t *testing.T) {
	type fields struct {
		s *Service
	}
	tests := []struct {
		name   string
		fields fields
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.NetImportSync(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.NetImportSync() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandService_MissingMoviesSearch(t *testing.T) {
	type fields struct {
		s *Service
	}
	type args struct {
		filter *Filter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandService{
				s: tt.fields.s,
			}
			if got := c.MissingMoviesSearch(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandService.MissingMoviesSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
