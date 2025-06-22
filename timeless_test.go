package timeless

import (
	"testing"
	"time"
)

func TestTimezoneToOffset(t *testing.T) {
	tests := map[string]UTCOffset{
		"UTC+1":  1,
		"UTC-5":  -5,
		"UTC+11": 11,
		"UTC":    0,
		"EST":    -5,
	}

	for s, n := range tests {
		if o := TimezoneToOffset(s); o != n {
			t.Fatalf("%s should return %d but returns %d", s, n, o)
		}
	}
}

func TestParseTimeLength(t *testing.T) {
	tests := map[string]time.Duration{
		"1h2m":                time.Hour + time.Minute*1,
		"2d 8h 30s":           Day*2 + time.Hour*8 + time.Second*30,
		"5weeks1day 7minutes": Week*5 + Day + time.Minute*5,
	}
	for s, n := range tests {
		if o := ParseTimeLength(s); o != n {
			t.Fatalf("%s should return %d but returns %d", s, n, o)
		}
	}
}
