package cmd

import (
	"errors"
	"image"
	"log"

	kabanIo "github.com/aethiopicuschan/kaban/cmd/io"
	"github.com/aethiopicuschan/kaban/merge"

	"github.com/spf13/cobra"
)

// packCmd represents the pack command.
var packCmd = &cobra.Command{
	Use:   "pack [src.png] [flags]",
	Short: "Pack sprite sheet image",
	Long:  "Pack sprite sheet image",
	Args:  cobra.MinimumNArgs(1),
	RunE:  pack,
}

func initPackCmd() {
	packCmd.Flags().StringP("output", "o", "./packed.png", "output file")
	packCmd.Flags().StringP("mode", "m", "grid", `packing mode. allowed: "grid"`)
}

func pack(c *cobra.Command, args []string) (err error) {
	output, err := c.Flags().GetString("output")
	if err != nil {
		return
	}
	mode, err := c.Flags().GetString("mode")
	if err != nil {
		return
	}
	if mode != "grid" {
		return errors.New("invalid mode")
	}

	// Read source images.
	var imgs []image.Image
	for _, path := range args {
		img, err := kabanIo.ReadImage(path)
		if err != nil {
			return err
		}
		imgs = append(imgs, img)
	}

	// Merge images.
	mergedImg, err := merge.Merge(imgs...)
	if err != nil {
		return
	}

	// Write merged image.
	err = kabanIo.WriteImage(output, mergedImg)
	if err != nil {
		return
	}

	log.Printf("Packed %d images to %s!", len(args), output)
	return
}
