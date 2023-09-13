package io

import (
	"image"
	"image/png"
	"os"
)

func WriteImage(path string, img image.Image) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()
	err = png.Encode(file, img)
	return
}
