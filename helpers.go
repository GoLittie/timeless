package timeless

import (
	"strings"
	"time"
)

func isDigit(c uint8) bool {
	return '0' <= c && c <= '9'
}

func isWhitespace(c uint8) bool {
	return c == ' ' || c == '\n' || c == '\r'
}

func isDateSeparator(c uint8) bool {
	return c == '-' || c == '.' || c == '/' || c == '\\'
}

func periodToMS(period string) int {
	switch strings.ToLower(period) {
	case "s", "second", "seconds":
		return 1000
	case "m", "minute", "minutes":
		return 60 * 1000
	case "h", "hour", "hours":
		return 60 * 60 * 1000
	case "d", "day", "days":
		return 24 * 60 * 60 * 1000
	case "w", "week", "weeks":
		return 7 * 24 * 60 * 60 * 1000
	case "mo", "month", "months":
		return 30 * 24 * 60 * 60 * 1000
	case "y", "year", "years":
		return 365 * 24 * 60 * 60 * 1000
	default:
		return 0
	}
}

func reorderFormat(num1, num2, num3 int, format DateFormat) (year int, month int, day int) {
	if format > 2 || format < 0 {
		format = DEFAULT_DATE_FORMAT
	}
	switch format {
	case DDMMYY:
		return num3, num2, num1
	case MMDDYY:
		return num3, num1, num2
	case YYMMDD:
		return num1, num2, num3
	}
	return
}

func getLocalTimezoneOffset() UTCOffset {
	_, o := time.Now().Zone()
	return UTCOffset(o / 3600)
}
