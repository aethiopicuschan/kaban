package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "kaban",
	Long:              "Kaban is a simple tool for manipulating sprite sheet images.",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Version:           "0.0.3",
}

func init() {
	unpackCmd.Flags().StringP("output-dir", "o", "./", "output directory")
	// Size limit.
	unpackCmd.Flags().IntP("min-width", "", 0, "minimum width of sprite")
	unpackCmd.Flags().IntP("max-width", "", 0, "max width of sprite")
	unpackCmd.Flags().IntP("min-height", "", 0, "minimum height of sprite")
	unpackCmd.Flags().IntP("max-height", "", 0, "max height of sprite")
	rootCmd.AddCommand(unpackCmd)
	rootCmd.AddCommand(packCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
