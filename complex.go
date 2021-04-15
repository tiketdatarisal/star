package star

func Complex64(a, b float32) *complex64 {
	c := complex(a, b)
	return &c
}

func Complex128(a, b float64) *complex128 {
	c := complex(a, b)
	return &c
}
