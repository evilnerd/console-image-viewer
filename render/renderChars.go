package render

import (
	"console-image-previewer/processing"
	"fmt"
	"image"
	"math"
)

func RenderImageChars(img image.Image, imgType string) {

	newX, newY, resizedImg := processing.Resize(img, imgType)

	greyImage := processing.RgbaToGray(resizedImg)

	chars := " .:-=+*#%@"
	numChars := len(chars)
	section := (float64(numChars-1) / 255)

	for y := 0; y < newY; y++ {
		for x := 0; x < newX; x++ {
			//			col := greyImage.At(x, y)
			y := greyImage.GrayAt(x, y).Y
			i := int(math.Floor(section * float64(y)))
			c := chars[i : i+1]

			fmt.Print(c)
			fmt.Print(c)
		}
		fmt.Println()
	}
}
