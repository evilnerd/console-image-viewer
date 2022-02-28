package main

import (
	"console-image-previewer/cmd"
	log "github.com/pterm/pterm"
)

func main() {

	log.DefaultHeader.Println("CIV - Console Image Viewer")

	cmd.Execute()

}
