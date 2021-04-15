package star

import (
	"testing"
	"time"
)

func TestByte(t *testing.T) {
	want := byte(78)
	got := Byte(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestRune(t *testing.T) {
	want := '£'
	got := Rune(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestString(t *testing.T) {
	want := "7819"
	got := String(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestBool(t *testing.T) {
	want := true
	got := Bool(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestTime(t *testing.T) {
	want := time.Date(1978, 1, 1, 0, 0, 0, 0, time.Local)
	got := Time(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestDuration(t *testing.T) {
	want := 1978 * time.Second
	got := Duration(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}
