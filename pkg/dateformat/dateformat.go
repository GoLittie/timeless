package dateformat

var DefaultDateFormat = DDMMYY

type DateFormat int

const (
	DDMMYY DateFormat = iota
	YYMMDD
	MMDDYY
)
