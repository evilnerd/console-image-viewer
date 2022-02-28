package render

import (
	"console-image-previewer/processing"
	"fmt"
	"github.com/muesli/termenv"
	"image"
)

func RenderImageColor(img image.Image, imgType string) {

	newX, newY, resizedImg := processing.Resize(img, imgType)

	p := termenv.ColorProfile()

	for y := 0; y < newY; y++ {
		for x := 0; x < newX; x++ {
			col := resizedImg.At(x, y)
			out := termenv.String("  ").Background(p.FromColor(col))
			fmt.Print(out)
		}
		fmt.Println()
	}
}
