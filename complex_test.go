package star

import "testing"

func TestComplex64(t *testing.T) {
	want := complex(float32(78), float32(19))
	got := Complex64(float32(78), float32(19))
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestComplex128(t *testing.T) {
	want := complex(float64(78), float64(19))
	got := Complex128(float64(78), float64(19))
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}
