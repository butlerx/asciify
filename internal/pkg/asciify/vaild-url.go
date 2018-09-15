package asciify

import "net/url"

// IsValidURL checks if a url is valid
func isValidURL(toTest string) bool {
	if _, err := url.ParseRequestURI(toTest); err != nil {
		return false
	}
	return true
}
