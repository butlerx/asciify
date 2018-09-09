package ascii

import (
	"bytes"
	"image"
	"image/color"
	"reflect"
)

// Pixel2Char Converts pixel too grey scale and replaces with a character whose intensity is similar
func Pixel2Char(img image.Image, height, width int) []byte {
	chars := "MND8OZ$7I?+=~:,.."
	buf := new(bytes.Buffer)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(chars[pos])
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}
