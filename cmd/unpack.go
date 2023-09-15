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
	RunE:  unpack,
}

func initUnpackCmd() {
	unpackCmd.Flags().StringP("output-dir", "o", "./", "output directory")
	unpackCmd.Flags().IntP("min-width", "", 0, "minimum width of sprite")
	unpackCmd.Flags().IntP("max-width", "", 0, "max width of sprite")
	unpackCmd.Flags().IntP("min-height", "", 0, "minimum height of sprite")
	unpackCmd.Flags().IntP("max-height", "", 0, "max height of sprite")
}

// Create functional options from flags.
func optionsUnpackCmd(c *cobra.Command) (opts []func(*detection.Option), err error) {
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

func unpack(c *cobra.Command, args []string) (err error) {
	// Check output directory.
	outputDir, err := c.Flags().GetString("output-dir")
	if err != nil {
		return err
	}
	exist := kabanIo.IsDirExist(outputDir)
	if !exist {
		return fmt.Errorf(`output directory "%s" does not exist`, outputDir)
	}
	src := args[0]
	// Read image.
	img, err := kabanIo.ReadImage(src)
	if err != nil {
		return
	}
	// Create functional options from flags.
	opts, err := optionsUnpackCmd(c)
	if err != nil {
		return
	}
	// Detect rects.
	rects, err := detection.Detect(img, opts...)
	if err != nil {
		return err
	}
	// Write cropped images.
	for _, rect := range rects {
		croppedImage := crop(img, rect)
		filename := fmt.Sprintf("%d_%d__%d_%d.png", rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
		pathToWrite := filepath.Join(outputDir, filename)
		if err := kabanIo.WriteImage(pathToWrite, croppedImage); err != nil {
			return err
		}
	}
	log.Printf("Unpacked %d images!", len(rects))
	return
}
