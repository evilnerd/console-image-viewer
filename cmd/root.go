package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "civ",
	Short: "CIV - Console Image Viewer renders images in text form on the console.",
	Long: `Use this to preview images from the comfort of your console, or perhaps over an SSH connection. 
Repo can be found at https://github.com/evilnerd/console-image-viewer`,
}

func init() {
	cobra.MinimumNArgs(1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
