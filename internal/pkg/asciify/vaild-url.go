package asciify

import (
	"net/url"
	"testing"
)

// IsValidURL checks if a url is valid.
func isValidURL(toTest string) bool {
	if _, err := url.ParseRequestURI(toTest); err != nil {
		return false
	}

	return true
}

// TestIsValidURL tests function isValidURL.
func TestIsValidURL(t *testing.T) {
	validURL := isValidURL("http://example.com/test?arg=1&arg=2")
	if !validURL {
		t.Errorf("isValidURL(\"http://example.com/test?arg=1&arg=2\") = %v; want true", validURL)
	}

	invalidURL := isValidURL("Not a url")
	if !invalidURL {
		t.Errorf("isValidURL(\"Not a url\") = %v; want false", invalidURL)
	}
}
