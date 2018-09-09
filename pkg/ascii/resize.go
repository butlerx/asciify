package ascii

import (
	"image"

	"github.com/nfnt/resize"
)

// Resize resizes the image into the final width while maintaining aspect ratio
func Resize(img image.Image, width int) (image.Image, int, int) {
	sz := img.Bounds()
	height := (sz.Max.Y * width * 10) / (sz.Max.X * 16)
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3), width, height
}
