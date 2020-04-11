package radarr

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	internal "github.com/SkYNewZ/radarr/internal/radarr"
)

func Test_newHistoryService(t *testing.T) {
	s := &Service{client: http.DefaultClient, url: internal.DummyURL}
	tests := []struct {
		name    string
		service *Service
		want    *HistoryService
	}{
		{
			name:    "Constructor",
			service: s,
			want:    &HistoryService{s},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHistoryService(tt.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHistoryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistoryService_Get(t *testing.T) {
	goodService := newHistoryService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
	})

	var expectedResponse *History
	err := json.Unmarshal([]byte(internal.DummyHistoryResponse), &expectedResponse)
	if err != nil {
		t.Fatal("Error json.Unmarshal()")
	}

	tests := []struct {
		name    string
		service *Service
		want    *Records
		wantErr bool
	}{
		{
			name:    "Should return expected response",
			service: goodService.s,
			want:    &expectedResponse.Records,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &HistoryService{
				s: tt.service,
			}
			got, err := s.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("HistoryService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HistoryService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistoryService_paginate(t *testing.T) {
	goodService := newHistoryService(&Service{
		client: internal.DummyHTTPClient,
		url:    internal.DummyURL,
	})

	var expectedResponse *History
	err := json.Unmarshal([]byte(internal.DummyHistoryResponse), &expectedResponse)
	if err != nil {
		t.Fatal("Error json.Unmarshal()")
	}

	type args struct {
		page int
	}

	tests := []struct {
		name    string
		service *Service
		args    args
		want    *History
		wantErr bool
	}{
		{
			name:    "Should return expected response",
			args:    args{1},
			service: goodService.s,
			want:    expectedResponse,
			wantErr: false,
		},
		{
			name:    "Should return 404",
			args:    args{2},
			service: goodService.s,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Should JSON error",
			args:    args{3},
			service: goodService.s,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Server error",
			args:    args{4},
			service: goodService.s,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &HistoryService{
				s: tt.service,
			}
			got, err := s.paginate(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("HistoryService.paginate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HistoryService.paginate() = %v, want %v", got, tt.want)
			}
		})
	}
}
