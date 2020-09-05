package ascii

import (
	"image"

	"github.com/nfnt/resize"
)

// Resize resizes the image into the final width while maintaining aspect ratio.
func Resize(img image.Image, width int) image.Image {
	sz := img.Bounds()

	return resize.Resize(
		uint(width),
		uint((sz.Max.Y*width*10)/(sz.Max.X*16)),
		img,
		resize.Lanczos3,
	)
}
