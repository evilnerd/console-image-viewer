package cmd

import (
	"console-image-previewer/input"
	"console-image-previewer/render"
	log "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	blocksCmd = &cobra.Command{
		Use:   "blocks <file or pattern1> [file or pattern2]...",
		Short: "Renders the picture in color or b/w blocks",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			matches := input.GetImageFilenames(args)
			for _, file := range matches {
				log.DefaultSection.Println(file)
				img, imgType := input.GetImage(file)
				render.ImageBlock(img, activeMaxEdge(), imgType, mono)
			}
		},
	}
	mono bool
)

func init() {
	blocksCmd.Flags().BoolVarP(&mono, "mono", "m", false, "Show picture in monochome (grayscale)")
	rootCmd.AddCommand(blocksCmd)
}
