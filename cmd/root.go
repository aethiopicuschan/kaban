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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
