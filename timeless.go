package timeless

import (
	"github.com/golittie/timeless/internal/lexer"
	"github.com/golittie/timeless/pkg/dateformat"
	time_calculator "github.com/golittie/timeless/pkg/time-calculator"
	"github.com/golittie/timeless/pkg/timezone"
	"strings"
	"time"
)

// Parse parses a string containing dates, times and lengths of time (5m/5minutes) to a time object.
func Parse(timeString string, opts ...ParseOption) time.Time {
	options := parseOptions{
		UTCOffset: timezone.SystemTimezone, DateFormat: dateformat.DefaultDateFormat,
	}

	for _, opt := range opts {
		opt(&options)
	}

	timeCalc := time_calculator.NewTimeCalculator(options.UTCOffset)

	l := lexer.NewLexer(timeString, options.DateFormat, options.AllowNegatives)
	for {
		statement := l.NextStatement()
		if statement == nil {
			break
		}

		switch statement.Type() {
		case lexer.TimelengthStatementType:
			timelength := statement.(lexer.TimelengthStatement)
			timeCalc.AddPeriod(timelength.Value, timelength.Period)
		case lexer.DateStatementType:
			date := statement.(lexer.DateStatement)
			timeCalc.SetDate(date.Year, date.Month, date.Day)
		case lexer.DaytimeStatementType:
			daytime := statement.(lexer.DaytimeStatement)
			timeCalc.SetDayTime(daytime.Hour, daytime.Minute, daytime.Second)
		}
	}

	return timeCalc.Calc()
}

// ParseRelativeTimeLength parses a string with time lengths relative to time.now().
func ParseRelativeTimeLength(timeLength string, opts ...ParseOption) time.Time {
	options := parseOptions{}

	for _, opt := range opts {
		opt(&options)
	}

	// as dates and daytimes aren't parsed: timezone doesn't matter
	timeCalc := time_calculator.NewTimeCalculator(0)

	l := lexer.NewLexer(timeLength, 0, options.AllowNegatives)
	for {
		statement := l.NextStatement()
		if statement == nil {
			break
		}

		switch statement.Type() {
		case lexer.TimelengthStatementType:
			timelength := statement.(lexer.TimelengthStatement)
			timeCalc.AddPeriod(timelength.Value, timelength.Period)
		}
	}

	return timeCalc.Calc()
}

func ParseDate(dateString string, opts ...ParseOption) time.Time {
	options := parseOptions{
		UTCOffset: timezone.SystemTimezone, DateFormat: dateformat.DefaultDateFormat,
	}

	for _, opt := range opts {
		opt(&options)
	}

	timeCalc := time_calculator.NewTimeCalculator(options.UTCOffset)

	l := lexer.NewLexer(dateString, options.DateFormat, options.AllowNegatives)
	for {
		statement := l.NextStatement()
		if statement == nil {
			break
		}

		switch statement.Type() {
		case lexer.DateStatementType:
			date := statement.(lexer.DateStatement)
			timeCalc.SetDate(date.Year, date.Month, date.Day)
		case lexer.DaytimeStatementType:
			daytime := statement.(lexer.DaytimeStatement)
			timeCalc.SetDayTime(daytime.Hour, daytime.Minute, daytime.Second)
		}
	}

	return timeCalc.Calc()
}

// ParseTimeLength parses a string of time lengths to a time.Duration approximation. Unlike the other functions this won't be accurate.
func ParseTimeLength(timeLength string, opts ...ParseOption) (duration time.Duration) {
	options := parseOptions{
		AllowNegatives: true,
	}

	for _, opt := range opts {
		opt(&options)
	}

	l := lexer.NewLexer(timeLength, 0, options.AllowNegatives)
	for {
		statement := l.NextStatement()
		if statement == nil {
			break
		}

		switch statement.Type() {
		case lexer.TimelengthStatementType:
			timelength := statement.(lexer.TimelengthStatement)
			duration += time.Duration(timelength.Value*periodToSecs(timelength.Period)) * time.Second
		}
	}

	return
}

func periodToSecs(period string) int {
	switch strings.ToLower(period) {
	case "s", "second", "seconds":
		return 1
	case "m", "minute", "minutes":
		return 60
	case "h", "hour", "hours":
		return 60 * 60
	case "d", "day", "days":
		return 24 * 60 * 60
	case "w", "week", "weeks":
		return 7 * 24 * 60 * 60
	case "mo", "month", "months":
		return 30 * 24 * 60 * 60
	case "y", "year", "years":
		return 365 * 24 * 60 * 60
	default:
		return 0
	}
}
