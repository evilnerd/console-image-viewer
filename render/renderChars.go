package render

import (
	"console-image-previewer/processing"
	"fmt"
	"image"
	"math"
)

const (
	CHARS = " .:-=+*#%@"
)

func ImageChars(img image.Image, maxEdge int, imgType string, invert bool) {

	newX, newY, resizedImg := processing.Resize(img, maxEdge, imgType)

	greyImage := processing.RgbaToGray(resizedImg)

	chars := CHARS

	if invert {
		chars = Reverse(chars)
	}

	numChars := len(chars)
	section := float64(numChars-1) / 255

	for y := 0; y < newY; y++ {
		for x := 0; x < newX; x++ {
			y := greyImage.GrayAt(x, y).Y
			i := int(math.Floor(section * float64(y)))
			c := chars[i : i+1]

			fmt.Print(c)
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
