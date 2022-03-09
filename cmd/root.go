package cmd

import (
	"console-image-previewer/processing"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "civ",
		Short: "CIV - Console Image Viewer renders images in text form on the console.",
		Long: `Use this to preview images from the comfort of your console, or perhaps over an SSH connection. 
Repo can be found at https://github.com/evilnerd/console-image-viewer`,
	}
	maxEdge = -1
)

const (
	MAXEDGE = "maxedge"
)

func init() {
	cobra.MinimumNArgs(1)
	rootCmd.PersistentFlags().IntVarP(&maxEdge, MAXEDGE, "e", 48, "Sets the max size to render the image with. The orientation of the image determines whether this applies to the width or height.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func activeMaxEdge() int {
	if rootCmd.PersistentFlags().Changed(MAXEDGE) {
		return maxEdge
	}
	return processing.DetermineMaxEdge()
}
