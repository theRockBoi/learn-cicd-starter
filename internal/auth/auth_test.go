package auth

import (
	"testing"
	"errors"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name   string
		header http.Header
		want   string
		err    error
	}{
		{"invalid api key", http.Header{"Authorization": []string{""}}, "", ErrNoAuthHeaderIncluded},
		{"valid api key", http.Header{"Authorization": []string{"ApiKey my-secret-key"}}, "my-secret-key", nil},
		{"malformed header", http.Header{"Authorization": []string{"Bearer my-secret-key"}}, "", errors.New("malformed authorization header")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.header)
			if (err != nil) != (tt.err != nil) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}