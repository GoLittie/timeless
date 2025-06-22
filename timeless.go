package timeless

import (
	"strconv"
	"strings"
	"time"
)

var DEFAULT_DATE_FORMAT = DDMMYY
var SYSTEM_TIMEZONE = getLocalTimezoneOffset()

func TimezoneToOffset(timezoneString string) UTCOffset {
	if len(timezoneString) > 4 && (timezoneString[:4] == "UTC-" || timezoneString[:4] == "UTC+") {
		s := timezoneString[3:]
		n, _ := strconv.Atoi(s)
		return UTCOffset(n)
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

// Parse parses a string containing dates, times and lengths of time (5m/5minutes) to a time object.
func Parse(timeString string, opts ...ParseOption) time.Time {
	options := parseOptions{
		UTCOffset: SYSTEM_TIMEZONE, DateFormat: DEFAULT_DATE_FORMAT,
	}

	for _, opt := range opts {
		opt(&options)
	}

	var extraMs int64
	l := newLexer(timeString)
	dateD := dateData{0, 0, 0}
	timeD := timeData{0, 0, 0}
	for !l.atEnd(0) {
		num := l.nextNumber(true)
		if num == "" {
			l.c++
			continue
		}
		nc := l.nextChars()

		if isDateSeparator(nc[0]) { // Parse date
			if strings.HasPrefix(num, "-") {
				num = num[1:]
			}

			n1, _ := strconv.Atoi(num)
			n2, _ := strconv.Atoi(l.nextNumber(false))
			var n3 int
			if !l.atEnd(0) && isDateSeparator(l.s[l.c]) {
				l.c++
				n3, _ = strconv.Atoi(l.nextNumber(false))
			}

			dateD.year, dateD.month, dateD.day = reorderFormat(n1, n2, n3, options.DateFormat)
		} else if nc == ":" { // Parse time
			if strings.HasPrefix(num, "-") {
				num = num[1:]
			}

			// Set hours
			timeD.hour, _ = strconv.Atoi(num)
			timeD.hour %= 24
			// Set minutes
			num = l.nextNumber(false)
			timeD.minute, _ = strconv.Atoi(num)
			timeD.minute %= 60

			if !l.atEnd(0) && l.s[l.c] == ':' {
				// Set seconds
				l.c++
				num = l.nextNumber(false)
				timeD.second, _ = strconv.Atoi(num)
				timeD.second %= 60
			}
		} else if nc != "" { // Parse time length
			p := periodToMS(nc)
			if p != 0 {
				n, _ := strconv.Atoi(num)
				extraMs += int64(n * p)
			}
		}

	}

	var t time.Time
	if dateD.day != 0 { // Time from date
		t = time.Date(dateD.year, time.Month(dateD.month), dateD.day, timeD.hour, timeD.minute, timeD.second, 0, time.FixedZone("", int(options.UTCOffset*60*60)))
	} else { // Time from now
		t = time.Now()
		if timeD.hour != 0 {
			now := time.Now()
			t = t.Add(time.Duration(timeD.hour-now.Hour()) * time.Hour).
				Add(time.Duration(timeD.minute-now.Minute()) * time.Minute).
				Add(time.Duration(timeD.second-now.Second()) * time.Second)
		}
	}
	t = t.Add(time.Duration(extraMs) * time.Millisecond).
		Add(time.Duration(options.UTCOffset-SYSTEM_TIMEZONE) * time.Hour)
	return t
}

// ParseTimeLength parses a string with time lengths to number of milliseconds.
func ParseTimeLength(timeLength string, opts ...ParseOption) (duration time.Duration) {
	options := parseOptions{
		AllowNegatives: true,
	}

	for _, opt := range opts {
		opt(&options)
	}

	l := newLexer(timeLength)

	for !l.atEnd(0) {
		num := l.nextNumber(options.AllowNegatives)
		if num == "" {
			l.c++
			continue
		}
		nc := l.nextChars()
		if nc == "" {
			l.c++
			continue
		}
		p := periodToMS(nc)
		if p != 0 {
			n, _ := strconv.Atoi(num)
			duration += time.Duration(n*p) * time.Millisecond
		}
	}

	return
}

// ParseRelativeTimeLength parses a string with time lengths to a time object based on now.
func ParseRelativeTimeLength(timeLength string, opts ...ParseOption) time.Time {
	options := parseOptions{
		UTCOffset: SYSTEM_TIMEZONE,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return time.Now().
		Add(ParseTimeLength(timeLength)).
		Add(time.Hour * time.Duration(options.UTCOffset-SYSTEM_TIMEZONE))
}

func WithTimezone(offset UTCOffset) ParseOption {
	return func(options *parseOptions) {
		options.UTCOffset = offset
	}
}

func WithDateFormat(dateFormat DateFormat) ParseOption {
	return func(options *parseOptions) {
		options.DateFormat = dateFormat
	}
}

func WithoutNegatives() ParseOption {
	return func(options *parseOptions) {
		options.AllowNegatives = false
	}
}
