package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test_api_key")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if apiKey != "test_api_key" {
		t.Errorf("Expected API key 'test_api_key', got %v", apiKey)
	}
}

func TestGetAPIKey_Errors(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantError error
	}{
		{"No Authorization header", http.Header{}, ErrNoAuthHeaderIncluded},
		{"Malformed Authorization header", http.Header{"Authorization": {"Api test_api_key"}}, errors.New("malformed authorization header")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetAPIKey(tt.headers)
			if err == nil || err.Error() != tt.wantError.Error() {
				t.Errorf("Expected error %v, got %v", tt.wantError, err)
			}
		})
	}
}
