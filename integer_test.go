package star

import "testing"

func TestInt(t *testing.T) {
	want := 78
	got := Int(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestInt8(t *testing.T) {
	want := int8(78)
	got := Int8(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestInt16(t *testing.T) {
	want := int16(78)
	got := Int16(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestInt32(t *testing.T) {
	want := int32(78)
	got := Int32(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestInt64(t *testing.T) {
	want := int64(78)
	got := Int64(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestUint(t *testing.T) {
	want := uint(78)
	got := Uint(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestUint8(t *testing.T) {
	want := uint8(78)
	got := Uint8(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestUint16(t *testing.T) {
	want := uint16(78)
	got := Uint16(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestUint32(t *testing.T) {
	want := uint32(78)
	got := Uint32(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}

func TestUint64(t *testing.T) {
	want := uint64(78)
	got := Uint64(want)
	if *got != want {
		t.Errorf("Test failed: Want %v, got %v", want, *got)
	}
}
