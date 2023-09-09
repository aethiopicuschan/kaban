package cmd

import (
	"fmt"
	"image"
	"log"

	kabanIo "github.com/aethiopicsuchan/kaban/cmd/io"
	"github.com/aethiopicsuchan/kaban/detection"
	"github.com/spf13/cobra"
)

// unpackCmd represents the unpack command.
var unpackCmd = &cobra.Command{
	Use:  "unpack <src.png> [flags]",
	Long: "Unpack sprite sheet image.",
	Args: cobra.ExactArgs(1),
	Run:  main,
}

func init() {
	unpackCmd.Flags().StringP("output-dir", "o", "./", "output directory")
	// Size limit.
	unpackCmd.Flags().IntP("min-width", "", 0, "minimum width of sprite")
	unpackCmd.Flags().IntP("max-width", "", 0, "max width of sprite")
	unpackCmd.Flags().IntP("min-height", "", 0, "minimum height of sprite")
	unpackCmd.Flags().IntP("max-height", "", 0, "max height of sprite")
	rootCmd.AddCommand(unpackCmd)
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

func main(c *cobra.Command, args []string) {
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
	exist = kabanIo.IsFileExist(src)
	if !exist {
		log.Fatalf("source file %s does not exist", src)
	}
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
		if err := kabanIo.WriteImage(outputDir, filename, croppedImage); err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("Unpacked %d images!", len(rects))
}
