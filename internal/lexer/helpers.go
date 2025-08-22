package lexer

import (
	"github.com/golittie/timeless/pkg/dateformat"
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

func reorderFormat(num1, num2, num3 int, format dateformat.DateFormat) (year int, month int, day int) {
	if format > 2 || format < 0 {
		format = dateformat.DefaultDateFormat
	}
	switch format {
	case dateformat.DDMMYY:
		return num3, num2, num1
	case dateformat.MMDDYY:
		return num3, num1, num2
	case dateformat.YYMMDD:
		return num1, num2, num3
	}
	return
}
