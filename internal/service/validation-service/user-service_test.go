package user_service

import (
	"net/http"
	"testing"
)

func TestIsValidHttpRequest(t *testing.T) {
	tests := []struct {
		name           string
		actualMethod   string
		expectedMethod string
		shouldFail     bool
	}{
		{"Valid GET", http.MethodGet, http.MethodGet, false},
		{"Valid POST", http.MethodPost, http.MethodPost, false},
		{"Valid PUT", http.MethodPut, http.MethodPut, false},
		{"Valid DELETE", http.MethodDelete, http.MethodDelete, false},
		{"Invalid mismatch", http.MethodPost, http.MethodGet, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.actualMethod, "http://localhost:8080", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			err = IsValidHttpRequest(req, tt.expectedMethod)
			if tt.shouldFail && err == nil {
				t.Errorf("Expected failure but got none")
			}
			if !tt.shouldFail && err != nil {
				t.Errorf("Expected success but got error: %v", err)
			}
		})
	}
}
