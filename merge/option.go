package merge

type Option struct {
	mode string
}

func WithGridMode() func(*Option) {
	return func(o *Option) {
		o.mode = "grid"
	}
}
