package detection

import (
	"image"
	"image/color"

	"github.com/aethiopicuschan/kaban/types"
)

func isTransparent(c color.Color) bool {
	_, _, _, a := c.RGBA()
	return a == 0
}

func growRegion(img image.Image, start types.Point, visited map[types.Point]bool) image.Rectangle {
	queue := []types.Point{start}
	minX, minY, maxX, maxY := start.X, start.Y, start.X, start.Y

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if _, ok := visited[p]; ok {
			continue
		}
		visited[p] = true

		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}

		for _, offset := range []types.Point{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1}} {
			neighbor := types.Point{X: p.X + offset.X, Y: p.Y + offset.Y}
			if neighbor.X >= 0 && neighbor.X < img.Bounds().Dx() &&
				neighbor.Y >= 0 && neighbor.Y < img.Bounds().Dy() &&
				!isTransparent(img.At(neighbor.X, neighbor.Y)) {
				queue = append(queue, neighbor)
			}
		}
	}

	return image.Rect(minX, minY, maxX+1, maxY+1)
}

// Detect rects from img.
func Detect(img image.Image, options ...func(*Option)) (rects []image.Rectangle, err error) {
	srcWidth := img.Bounds().Dx()
	srcHeight := img.Bounds().Dy()

	// Default options.
	opt := &Option{
		minWidth:  0,
		minHeight: 0,
		maxWidth:  srcWidth,
		maxHeight: srcHeight,
	}
	// Apply options.
	for _, o := range options {
		o(opt)
	}

	// Map of visited point.
	visited := make(map[types.Point]bool)

	for y := 0; y < srcHeight; y++ {
		for x := 0; x < srcWidth; x++ {
			if _, ok := visited[types.Point{X: x, Y: y}]; !ok && !isTransparent(img.At(x, y)) {
				rect := growRegion(img, types.Point{X: x, Y: y}, visited)
				// Ignore small rect in accordance with Option.
				width := rect.Max.X - rect.Min.X
				height := rect.Max.Y - rect.Min.Y
				if width < opt.minWidth || height < opt.minHeight || width > opt.maxWidth || height > opt.maxHeight {
					continue
				}
				rects = append(rects, rect)
			}
		}
	}

	return
}
