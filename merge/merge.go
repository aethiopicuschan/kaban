package merge

import (
	"errors"
	"image"
)

func Merge(imgs []image.Image, options ...func(*Option)) (mergedImg image.Image, m map[int]image.Point, err error) {
	// Default options.
	opt := &Option{
		mode: "grid",
	}
	// Apply options.
	for _, o := range options {
		o(opt)
	}

	err = errors.New("not implemented")
	return
}
