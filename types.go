package timeless

import "time"

type UTCOffset float32

const (
	EST  UTCOffset = -5
	UTC  UTCOffset = 0
	GMT  UTCOffset = 0
	BST  UTCOffset = 1
	CET  UTCOffset = 1
	CEST UTCOffset = 2
	IST  UTCOffset = 5.5
)

type DateFormat int

const (
	DDMMYY DateFormat = iota
	YYMMDD
	MMDDYY
)

type dateData struct {
	year  int
	month int
	day   int
}

type timeData struct {
	hour   int
	minute int
	second int
}

type parseOptions struct {
	UTCOffset
	DateFormat
	AllowNegatives bool
}

type ParseOption func(*parseOptions)

const Day = time.Hour * 24
const Week = Day * 7

// An oversimplification for convenience

const Month = Day * 30
const Year = Day * 365
