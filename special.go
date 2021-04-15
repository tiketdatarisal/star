package star

import "time"

func Byte(v byte) *byte {
	return &v
}

func Rune(v rune) *rune {
	return &v
}

func String(v string) *string {
	return &v
}

func Bool(v bool) *bool {
	return &v
}

func Time(v time.Time) *time.Time {
	return &v
}

func Duration(v time.Duration) *time.Duration {
	return &v
}
