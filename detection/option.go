package detection

type Option struct {
	minWidth  int
	minHeight int
	maxWidth  int
	maxHeight int
}

func WithMinWidth(minWidth int) func(*Option) {
	return func(o *Option) {
		o.minWidth = minWidth
	}
}

func WithMinHeight(minHeight int) func(*Option) {
	return func(o *Option) {
		o.minHeight = minHeight
	}
}

func WithMaxWidth(maxWidth int) func(*Option) {
	return func(o *Option) {
		o.maxWidth = maxWidth
	}
}

func WithMaxHeight(maxHeight int) func(*Option) {
	return func(o *Option) {
		o.maxHeight = maxHeight
	}
}
