package cmd

import (
	"fmt"
	"image"
	"log"
	"path/filepath"

	kabanIo "github.com/aethiopicuschan/kaban/cmd/io"
	"github.com/aethiopicuschan/kaban/detection"
	"github.com/spf13/cobra"
)

// unpackCmd represents the unpack command.
var unpackCmd = &cobra.Command{
	Use:   "unpack <src.png> [flags]",
	Short: "Unpack sprite sheet image",
	Long:  "Unpack sprite sheet image",
	Args:  cobra.ExactArgs(1),
	Run:   unpack,
}

// Create functional options from flags.
func options(c *cobra.Command) (opts []func(*detection.Option), err error) {
	funcMap := map[string]func(int) func(*detection.Option){
		"min-width":  detection.WithMinWidth,
		"max-width":  detection.WithMaxWidth,
		"min-height": detection.WithMinHeight,
		"max-height": detection.WithMaxHeight,
	}

	for name, f := range funcMap {
		if c.Flags().Changed(name) {
			value, err := c.Flags().GetInt(name)
			if err != nil {
				return nil, err
			}
			opts = append(opts, f(value))
			log.Printf("Set %s to %d", name, value)
		}
	}

	return
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func crop(img image.Image, rect image.Rectangle) image.Image {
	return img.(SubImager).SubImage(rect)
}

func unpack(c *cobra.Command, args []string) {
	// Check output directory.
	outputDir, err := c.Flags().GetString("output-dir")
	if err != nil {
		log.Fatal(err)
	}
	exist := kabanIo.IsDirExist(outputDir)
	if !exist {
		log.Fatalf("output directory %s does not exist", outputDir)
	}
	// Check source file.
	if len(args) != 1 {
		log.Fatalf("accepts 1 arg(s), received %d", len(args))
	}
	src := args[0]
	// Read image.
	img, err := kabanIo.ReadImage(src)
	if err != nil {
		log.Fatal(err)
	}
	// Create functional options from flags.
	opts, err := options(c)
	if err != nil {
		log.Fatal(err)
	}
	// Detect rects.
	rects, err := detection.Detect(img, opts...)
	if err != nil {
		log.Fatal(err)
	}
	// Write cropped images.
	for _, rect := range rects {
		croppedImage := crop(img, rect)
		filename := fmt.Sprintf("%d_%d__%d_%d.png", rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
		pathToWrite := filepath.Join(outputDir, filename)
		if err := kabanIo.WriteImage(pathToWrite, croppedImage); err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("Unpacked %d images!", len(rects))
}
