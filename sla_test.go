package main

import (
	"testing"
	"time"
)

func TestFormatSLAPeriod(t *testing.T) {
	tests := []struct {
		name       string
		downtime   float64
		period     float64
		want       string
		wantZero   bool
		wantFormat string
	}{
		{
			name:     "1 hour in a day",
			downtime: float64(time.Hour.Seconds()),
			period:   float64(24 * time.Hour.Seconds()),
			want:     "95.8333 %",
		},
		{
			name:     "1 day in a week",
			downtime: float64(24 * time.Hour.Seconds()),
			period:   float64(7 * 24 * time.Hour.Seconds()),
			want:     "85.7143 %",
		},
		{
			name:     "exceeds period",
			downtime: float64(25 * time.Hour.Seconds()),
			period:   float64(24 * time.Hour.Seconds()),
			want:     "0 %",
			wantZero: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatSLAPeriod(tt.downtime, tt.period)
			if got != tt.want {
				t.Errorf("formatSLAPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}
