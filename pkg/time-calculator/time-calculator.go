package time_calculator

import (
	"strings"
	"time"
)

type TimeCalculator struct {
	fixedTime   *timeParts
	bonusYears  int
	bonusMonths int
	bonusDays   int
	bonusSecs   int
	offset      float32
}

func (c *TimeCalculator) location() *time.Location {
	return time.FixedZone("", int(c.offset*3600))
}

func (c *TimeCalculator) SetDate(year int, month int, day int) {
	if c.fixedTime == nil {
		c.fixedTime = &timeParts{}
	}
	c.fixedTime.year = &year
	c.fixedTime.month = &month
	c.fixedTime.day = &day
}

func (c *TimeCalculator) SetYear(year int) {
	if c.fixedTime == nil {
		c.fixedTime = &timeParts{}
	}
	c.fixedTime.year = &year
	if c.fixedTime.month == nil {
		one := 1
		c.fixedTime.month = &one
	}
	if c.fixedTime.day == nil {
		one := 1
		c.fixedTime.day = &one
	}
}

func (c *TimeCalculator) SetMonth(month int) {
	if c.fixedTime == nil {
		c.fixedTime = &timeParts{}
	}
	c.fixedTime.month = &month
	if c.fixedTime.day == nil {
		one := 1
		c.fixedTime.day = &one
	}
}

func (c *TimeCalculator) SetDay(day int) {
	if c.fixedTime == nil {
		c.fixedTime = &timeParts{}
	}
	c.fixedTime.day = &day
}

func (c *TimeCalculator) SetDayTime(hour int, minute int, second int) {
	if c.fixedTime == nil {
		c.fixedTime = &timeParts{}
	}
	c.fixedTime.hour = hour
	c.fixedTime.minute = minute
	c.fixedTime.second = second
}

func (c *TimeCalculator) AddYears(years int) {
	c.bonusYears += years
}

func (c *TimeCalculator) AddMonths(months int) {
	c.bonusMonths += months
}

func (c *TimeCalculator) AddWeeks(weeks int) {
	c.bonusDays += weeks * 7
}

func (c *TimeCalculator) AddDays(days int) {
	c.bonusDays += days
}

func (c *TimeCalculator) AddHours(hours int) {
	c.bonusSecs += hours * 3600
}

func (c *TimeCalculator) AddMinutes(mins int) {
	c.bonusSecs += mins * 60
}

func (c *TimeCalculator) AddSecs(secs int) {
	c.bonusSecs += secs
}

func (c *TimeCalculator) AddPeriod(value int, period string) {
	switch strings.ToLower(period) {
	case "seconds", "secs", "s":
		c.AddSecs(value)
	case "minutes", "mins", "m":
		c.AddMinutes(value)
	case "hours", "hrs", "h":
		c.AddHours(value)
	case "weeks", "w":
		c.AddWeeks(value)
	case "days", "d":
		c.AddDays(value)
	case "months", "mo":
		c.AddMonths(value)
	case "years", "yr", "y":
		c.AddYears(value)
	}
}

func (c *TimeCalculator) Calc() time.Time {
	var t time.Time
	if c.fixedTime == nil {
		t = time.Now()
	} else {
		t = c.fixedTime.toTime(c.location())
	}
	t = t.AddDate(c.bonusYears, c.bonusMonths, c.bonusDays)
	t = t.Add(time.Duration(c.bonusSecs) * time.Second)
	return t
}

func NewTimeCalculator(offset float32) TimeCalculator {
	return TimeCalculator{
		offset: offset,
	}
}
