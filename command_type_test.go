package radarr

import (
	"reflect"
	"testing"
)

func TestFilter_get(t *testing.T) {
	tests := []struct {
		name string
		f    Filter
		want *filter
	}{
		{
			name: "FilterByMonitored",
			f:    FilterByMonitored,
			want: &availableFilter[FilterByMonitored],
		},
		{
			name: "FilterByNonMonitored",
			f:    FilterByNonMonitored,
			want: &availableFilter[FilterByNonMonitored],
		},
		{
			name: "FilterAll",
			f:    FilterAll,
			want: &availableFilter[FilterAll],
		},
		{
			name: "FilterByStatusAndAvailable",
			f:    FilterByStatusAndAvailable,
			want: &availableFilter[FilterByStatusAndAvailable],
		},
		{
			name: "FilterByStatusAndReleased",
			f:    FilterByStatusAndReleased,
			want: &availableFilter[FilterByStatusAndReleased],
		},
		{
			name: "FilterByStatusAndInCinemas",
			f:    FilterByStatusAndInCinemas,
			want: &availableFilter[FilterByStatusAndInCinemas],
		},
		{
			name: "FilterByStatusAndAnnounced",
			f:    FilterByStatusAndAnnounced,
			want: &availableFilter[FilterByStatusAndAnnounced],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter.get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImportMode_get(t *testing.T) {
	tests := []struct {
		name string
		i    ImportMode
		want string
	}{
		{
			name: "Move",
			i:    Move,
			want: availableImportMode[Move],
		},
		{
			name: "Copy",
			i:    Copy,
			want: availableImportMode[Copy],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("ImportMode.get() = %v, want %v", got, tt.want)
			}
		})
	}
}
