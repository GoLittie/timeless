package time_calculator

import "time"

type timeParts struct {
	year   *int
	month  *int
	day    *int
	hour   int
	minute int
	second int
}

func (tp timeParts) toTime(loc *time.Location) time.Time {
	t := time.Now()
	var year int
	if tp.year == nil {
		year = t.Year()
	} else {
		year = *tp.year
	}

	var month time.Month
	if tp.month == nil || *tp.month < 1 {
		month = t.Month()
	} else {
		month = time.Month(*tp.month)
	}

	var day int
	if tp.day == nil || *tp.day < 1 {
		day = t.Day()
	} else {
		day = *tp.day
	}

	return time.Date(year, month, day, tp.hour, tp.minute, tp.second, 0, loc)
}
