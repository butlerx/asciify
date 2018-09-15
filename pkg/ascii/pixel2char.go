package ascii

import (
	"bytes"
	"image"
	"image/color"
	"reflect"
)

// Pixel2Char Converts pixel too grey scale and replaces with a character whose intensity is similar
func Pixel2Char(img image.Image) []byte {
	chars := "MND8OZ$7I?+=~:,  "
	buf := new(bytes.Buffer)
	sz := img.Bounds()
	for i := 0; i < sz.Max.Y; i++ {
		for j := 0; j < sz.Max.X; j++ {
			_ = buf.WriteByte(
				chars[int(
					reflect.ValueOf(
						color.GrayModel.Convert(
							img.At(j, i),
						),
					).FieldByName("Y").Uint()*16/255,
				)],
			)
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}
