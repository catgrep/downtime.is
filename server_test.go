package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		wantStatus  int
		wantBody    []string
		notWantBody []string
	}{
		{
			name:       "no downtime",
			path:       "/0",
			wantStatus: http.StatusOK,
			wantBody: []string{
				"SLA uptime in case of 0s downtime",
				"Daily reporting:</strong> 100 %",
				"Weekly reporting:</strong> 100 %",
				"Monthly reporting:</strong> 100 %",
				"Quarterly reporting:</strong> 100 %",
				"Yearly reporting:</strong> 100 %",
			},
		},
		{
			name:       "1 hour downtime",
			path:       "/1h",
			wantStatus: http.StatusOK,
			wantBody: []string{
				"SLA uptime in case of 1h downtime",
				"Daily reporting:</strong> 95.8333 %",
				"Weekly reporting:</strong> 99.4048 %",
				"Monthly reporting:</strong> 99.8631 %",
				"Quarterly reporting:</strong> 99.9544 %",
				"Yearly reporting:</strong> 99.9886 %",
			},
		},
		{
			name:       "1 day downtime",
			path:       "/1d",
			wantStatus: http.StatusOK,
			wantBody: []string{
				"SLA uptime in case of 1d downtime",
				"Daily reporting:</strong> 0 %",
				"Weekly reporting:</strong> 85.7143 %",
				"Monthly reporting:</strong> 96.7145 %<",
				"Quarterly reporting:</strong> 98.9048 %",
				"Yearly reporting:</strong> 99.7262 %",
			},
		},
		{
			name:       "1 day and 12 hours downtime",
			path:       "/1d12h",
			wantStatus: http.StatusOK,
			wantBody: []string{
				"SLA uptime in case of 1d 12h downtime",
				"Daily reporting:</strong> 0 %",
				"Weekly reporting:</strong> 78.5714 %",
				"Monthly reporting:</strong> 95.0718 %",
				"Quarterly reporting:</strong> 98.3573 %",
				"Yearly reporting:</strong> 99.5893 %",
			},
		},
		{
			name:       "complex duration",
			path:       "/5d12h34m30s",
			wantStatus: http.StatusOK,
			wantBody: []string{
				"SLA uptime in case of 5d 12h 34m 30s downtime",
				"Daily reporting:</strong> 0 %",
				"Weekly reporting:</strong> 21.0863 %",
				"Monthly reporting:</strong> 81.8511 %",
				"Quarterly reporting:</strong> 93.9504 %",
				"Yearly reporting:</strong> 98.4876 %",
			},
		},
		{
			name:       "invalid input",
			path:       "/invalid",
			wantStatus: http.StatusBadRequest,
			wantBody: []string{
				"Invalid input",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			handleRequest(w, req)

			if got := w.Code; got != tt.wantStatus {
				t.Errorf("handleRequest() status = %v, want %v", got, tt.wantStatus)
			}

			resp := w.Body.String()
			t.Logf("Response: %s", resp)
			for _, want := range tt.wantBody {
				if !strings.Contains(resp, want) {
					t.Errorf("handleRequest() body missing %q", want)
				}
			}

			for _, notWant := range tt.notWantBody {
				if strings.Contains(resp, notWant) {
					t.Errorf("handleRequest() body contains unwanted %q", notWant)
				}
			}
		})
	}
}
