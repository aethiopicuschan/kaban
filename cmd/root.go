package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "kaban",
	Long:              "Kaban is a simple tool for manipulating sprite sheet images.",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Version:           "0.1.0",
}

func init() {
	// Unpack command.
	initUnpackCmd()
	rootCmd.AddCommand(unpackCmd)

	// Pack command.
	initPackCmd()
	rootCmd.AddCommand(packCmd)
}

func Execute() (err error) {
	return rootCmd.Execute()
}
