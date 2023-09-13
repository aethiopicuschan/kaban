package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// packCmd represents the pack command.
var packCmd = &cobra.Command{
	Use:   "pack [src.png] [flags]",
	Short: "Pack sprite sheet image",
	Long:  "Pack sprite sheet image",
	Args:  cobra.MinimumNArgs(1),
	Run:   pack,
}

func pack(c *cobra.Command, args []string) {
	log.Fatal("Not implemented")
}
