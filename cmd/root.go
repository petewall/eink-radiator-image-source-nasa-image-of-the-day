package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const ImageTypeName = "nasa-image-of-the-day"

var rootCmd = &cobra.Command{
	Use:   ImageTypeName,
	Short: "Generate an image from NASA's image of the day",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
