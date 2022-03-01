package cmd

import (
	"console-image-previewer/input"
	"console-image-previewer/render"
	log "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var charsCmd = &cobra.Command{
	Use:   "chars <file or pattern1> [file or pattern2]...",
	Short: "Renders the picture in mono ascii characters",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		matches := input.GetImageFilenames(args)
		for _, file := range matches {
			log.DefaultSection.Println(file)
			img, imgType := input.GetImage(file)
			render.RenderImageChars(img, imgType)
		}
	},
}

func init() {
	rootCmd.AddCommand(charsCmd)
}
