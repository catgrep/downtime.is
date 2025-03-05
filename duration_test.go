package main

import (
	"testing"
	"time"
)

func TestParseDowntimeDuration(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    time.Duration
		wantErr bool
	}{
		{
			name: "60s",
			path: "60s",
			want: 60 * time.Second,
		},
		{
			name: "numeric seconds",
			path: "3600",
			want: time.Hour,
		},
		{
			name: "one hour",
			path: "1h",
			want: time.Hour,
		},
		{
			name: "one day",
			path: "1d",
			want: 24 * time.Hour,
		},
		{
			name: "complex duration with days",
			path: "1d12h",
			want: 36 * time.Hour,
		},
		{
			name: "complex duration",
			path: "1h30m",
			want: 90 * time.Minute,
		},
		{
			name: "with milliseconds",
			path: "1h30m500ms",
			want: 90*time.Minute + 500*time.Millisecond,
		},
		{
			name:    "invalid input",
			path:    "invalid",
			wantErr: true,
		},
		{
			name:    "invalid days format",
			path:    "1dx",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDowntimeDuration(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDowntimeDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("parseDowntimeDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
