package star

import (
	"testing"
	"time"
)

func TestPointerToInt(t *testing.T) {
	want := 78
	got := PointerTo(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestPointerToString(t *testing.T) {
	want := "Hello"
	got := PointerTo(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestPointerToFloat(t *testing.T) {
	want := 7.8
	got := PointerTo(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestPointerToTime(t *testing.T) {
	want := time.Now()
	got := PointerTo(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}
