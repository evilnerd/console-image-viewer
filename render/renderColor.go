package render

import (
	"console-image-previewer/processing"
	"fmt"
	"github.com/muesli/termenv"
	"image"
)

func ImageBlock(originalImg image.Image, maxEdge int, imgType string, mono bool) {

	newX, newY, resizedImg := processing.Resize(originalImg, maxEdge, imgType)
	var img image.Image

	img = resizedImg

	if mono {
		img = processing.RgbaToGray(img)
	}

	p := termenv.ColorProfile()

	for y := 0; y < newY; y++ {
		for x := 0; x < newX; x++ {
			col := img.At(x, y)
			out := termenv.String("  ").Background(p.FromColor(col))
			fmt.Print(out)
		}
		fmt.Println()
	}
}
