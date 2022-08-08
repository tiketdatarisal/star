package star

func PointerTo[T any](v T) *T {
	return &v
}
