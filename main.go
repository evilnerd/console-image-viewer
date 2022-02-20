package main

import (
	"fmt"
	"github.com/muesli/termenv"
	"github.com/nathan-fiscaletti/consolesize-go"
	log "github.com/pterm/pterm"
	"golang.org/x/image/draw"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"path/filepath"
)

var (
	maxEdge = 48
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal.Printf("Specify the file to show.")
	}

	maxEdge = determineMaxEdge()

	for _, arg := range os.Args[1:] {
		matches, err := filepath.Glob(arg)
		if err != nil {
			log.Error.Printf("Could not parse argument %s: %v\n", arg, err)
			continue
		}
		for _, file := range matches {
			log.DefaultSection.Println(file)
			renderImage(file)
		}
	}
}

func determineMaxEdge() int {
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

func renderImage(file string) {

	f, err := os.Open(file)
	if err != nil {
		log.Error.Printf("Could load image %v\n", err)
	}

	defer f.Close()

	img, imgType, err := image.Decode(f)
	if err != nil {
		log.Error.Printf("Could not decode image %v\n", err)
	}

	orgX, orgY := img.Bounds().Max.X, img.Bounds().Max.Y
	newX, newY := newSize(orgX, orgX, maxEdge)
	resizedImg := image.NewRGBA(image.Rect(0, 0, newX, newY))
	draw.NearestNeighbor.Scale(resizedImg, resizedImg.Rect, img, img.Bounds(), draw.Over, nil)

	log.Info.Printf("Decoded file of type %s\n", imgType)
	log.Info.Printf("Original dimensions = %dx%d, showing dimensions = %dx%d\n", orgX, orgY, newX, newY)

	//	greyImage := rgbaToGray(resizedImg)
	p := termenv.ColorProfile()

	//	termenv.ClearScreen()
	for y := 0; y < newY; y++ {
		for x := 0; x < newX; x++ {
			//			col := greyImage.At(x, y)
			col := resizedImg.At(x, y)

			out := termenv.String("  ").Background(p.FromColor(col))
			fmt.Print(out)
		}
		fmt.Println()
	}
}

func rgbaToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}
