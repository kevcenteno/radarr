package radarr

import (
	"testing"
)

func TestError_Error(t *testing.T) {
	type fields struct {
		Code    int
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		struct {
			name   string
			fields fields
			want   string
		}{
			fields: fields{Code: 1234, Message: "foo"},
			want:   "Radarr error: code 1234, message 'foo'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
