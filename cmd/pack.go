package cmd

import (
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
	Run:   pack,
}

func pack(c *cobra.Command, args []string) {
	output, err := c.Flags().GetString("output")
	if err != nil {
		log.Fatal(err)
	}

	// Read source images.
	var imgs []image.Image
	for _, path := range args {
		img, err := kabanIo.ReadImage(path)
		if err != nil {
			log.Fatal(err)
		}
		imgs = append(imgs, img)
	}

	// Merge images.
	mergedImg, err := merge.Merge(imgs...)
	if err != nil {
		log.Fatal(err)
	}

	// Write merged image.
	err = kabanIo.WriteImage(output, mergedImg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Packed %d images to %s!", len(args), output)
}
