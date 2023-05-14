package cmd

import (
	"console-image-previewer/input"
	"console-image-previewer/render"
	log "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	charsCmd = &cobra.Command{
		Use:   "chars <file or pattern1> [file or pattern2]...",
		Short: "Renders the picture in mono ascii characters",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			matches := input.GetImageFilenames(args)
			for _, file := range matches {
				log.DefaultSection.Println(file)
				img, imgType := input.GetImage(file)
				render.ImageChars(img, activeMaxEdge(), imgType, invert)
			}
		},
	}
	invert bool
)

func init() {
	charsCmd.Flags().BoolVarP(&invert, "invert", "i", false, "Invert the picture (flip light-dark)")
	rootCmd.AddCommand(charsCmd)
}
