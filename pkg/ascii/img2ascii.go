package ascii

import (
	"image"
	"io"

	_ "image/jpeg"
	_ "image/png"
)

// Generate Converts any Image too an ascii image
func Generate(reader io.Reader, width int) (string, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}
	return string(Pixel2Char(Resize(img, width))), nil
}
