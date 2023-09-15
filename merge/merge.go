package merge

import (
	"errors"
	"image"
	"image/draw"
	"math"
)

func getMaxSize(imgs []image.Image) (width, height int) {
	for _, img := range imgs {
		if img.Bounds().Max.X > width {
			width = img.Bounds().Max.X
		}
		if img.Bounds().Max.Y > height {
			height = img.Bounds().Max.Y
		}
	}
	return
}

func getGridDimensions(totalImages int) (col, row int) {
	sqrt := int(math.Sqrt(float64(totalImages)))
	col = 1
	row = totalImages

	for x := sqrt; x > 0; x-- {
		if totalImages%x == 0 {
			col = x
			row = totalImages / x
			break
		}
	}

	if totalImages > 4 && col == 1 {
		return getGridDimensions(totalImages + 1)
	} else {
		return
	}
}

func Merge(imgs []image.Image, options ...func(*Option)) (mergedImg image.Image, points []image.Point, err error) {
	if len(imgs) == 0 {
		err = errors.New("no images")
		return
	}

	// Default options.
	opt := &Option{
		mode: "grid",
	}
	// Apply options.
	for _, o := range options {
		o(opt)
	}

	colNum, rowNum := getGridDimensions(len(imgs))
	gridWidth, gridHeight := getMaxSize(imgs)
	mergedImg = image.NewRGBA(image.Rect(0, 0, (gridWidth+1)*rowNum, (gridHeight+1)*colNum))
	index := 0
	for x := 0; x < rowNum; x++ {
		for y := 0; y < colNum; y++ {
			// Draw image.
			offsetX := x * (gridWidth + 1)
			offsetY := y * (gridHeight + 1)
			bounds := image.Rect(offsetX, offsetY, offsetX+gridWidth, offsetY+gridHeight)
			draw.Draw(mergedImg.(*image.RGBA), bounds, imgs[index], image.Point{0, 0}, draw.Src)
			points = append(points, image.Point{offsetX, offsetY})
			index++
			if index >= len(imgs) {
				break
			}
		}
	}

	return
}
