package timezone

import (
	"strconv"
	"strings"
	"time"
)

var SystemTimezone = getLocalTimezoneOffset()

const (
	EST  = -5
	UTC  = 0
	GMT  = 0
	BST  = 1
	CET  = 1
	CEST = 2
	IST  = 5.5
)

func TimezoneToOffset(timezoneString string) float32 {
	if len(timezoneString) > 4 && (timezoneString[:4] == "UTC-" || timezoneString[:4] == "UTC+") {
		s := timezoneString[3:]
		n, _ := strconv.Atoi(s)
		return float32(n)
	}
	switch strings.ToLower(timezoneString) {
	case "gmt", "bst":
		return 1
	case "est":
		return -5
	default:
		return UTC
	}
}

func getLocalTimezoneOffset() float32 {
	_, o := time.Now().Zone()
	return float32(o / 3600)
}
