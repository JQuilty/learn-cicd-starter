package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	want := "my-secret-key"
	if got != want {
		t.Errorf("expected key %q, got %q", want, got)
	}
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := http.Header{}

	got, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
	if got != "" {
		t.Errorf("expected empty key, got %q", got)
	}
}
