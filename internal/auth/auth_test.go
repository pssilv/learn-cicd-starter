package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    bool
	}{
		{
			name: "Valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			wantAPIKey: "valid_api_key",
			wantErr:    false,
		},
		{
			name: "Invalid API key",
			headers: http.Header{
				"Authorization": []string{"InvalidApiKey key"},
			},
			wantAPIKey: "",
			wantErr:    true,
		},
		{
			name:       "Missing authorization header",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotApiKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error := %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotApiKey != tt.wantAPIKey {
				t.Errorf("GetAPIKey() gotApiKey := %v, want %v", gotApiKey, tt.wantAPIKey)
			}
		})
	}
}
