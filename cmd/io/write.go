package io

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
)

func WriteImage(dir, name string, img image.Image) (err error) {
	pathToWrite := filepath.Join(dir, name)
	file, err := os.Create(pathToWrite)
	if err != nil {
		return
	}
	defer file.Close()
	err = png.Encode(file, img)
	return
}
