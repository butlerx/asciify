package asciify

import (
	"io"
	"net/http"
	"os"
)

// GetFile gets io reader for a remote ot local file
func GetFile(path string) (io.Reader, error) {
	if isValidURL(path) {
		response, err := http.Get(path)
		return response.Body, err
	}
	return os.Open(path)
}
