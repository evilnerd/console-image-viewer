package input

import (
	log "github.com/pterm/pterm"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func GetImage(fileName string) (image.Image, string) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Error.Printf("Could load image %v\n", err)
	}

	defer f.Close()

	img, imgType, err := image.Decode(f)
	if err != nil {
		log.Error.Printf("Could not decode image %v\n", err)
	}
	return img, imgType
}
