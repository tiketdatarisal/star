package star

import "testing"

func TestFloat32(t *testing.T) {
	want := float32(78.19)
	got := Float32(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestFloat64(t *testing.T) {
	want := 78.19
	got := Float64(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}
