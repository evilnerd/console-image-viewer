package processing

import (
	"github.com/nathan-fiscaletti/consolesize-go"
	log "github.com/pterm/pterm"
	"golang.org/x/image/draw"
	"image"
	"math"
)

func DetermineMaxEdge() int {
	width, height := consolesize.GetConsoleSize()

	if width > height {
		return height
	}
	return width

}

func newSize(x int, y int, maxEdge int) (newX int, newY int) {

	if x > y {
		newX = maxEdge
		newY = int(math.Floor((float64(y) / float64(x)) * float64(maxEdge)))
	} else {
		newX = int(math.Floor((float64(x) / float64(y)) * float64(maxEdge)))
		newY = maxEdge
	}
	return
}

func Resize(img image.Image, maxEdge int, imgType string) (int, int, *image.RGBA) {

	orgX, orgY := img.Bounds().Max.X, img.Bounds().Max.Y
	newX, newY := newSize(orgX, orgY, maxEdge)
	resizedImg := image.NewRGBA(image.Rect(0, 0, newX, newY))
	draw.NearestNeighbor.Scale(resizedImg, resizedImg.Rect, img, img.Bounds(), draw.Over, nil)

	log.Info.Printf("Decoded file of type %s\n", imgType)
	log.Info.Printf("Original dimensions = %dx%d, showing dimensions = %dx%d\n", orgX, orgY, newX, newY)
	return newX, newY, resizedImg
}
