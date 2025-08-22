package lexer

import (
	"github.com/golittie/timeless/pkg/dateformat"
	"strconv"
	"strings"
)

type Lexer struct {
	c          int
	s          string
	dateFormat dateformat.DateFormat
	negatives  bool
}

func (l *Lexer) atEnd(peek int) bool {
	return l.c+peek >= len(l.s)
}

func (l *Lexer) eatWhitespaces() {
	for !l.atEnd(0) && isWhitespace(l.s[l.c]) {
		l.c++
	}
}

func (l *Lexer) nextNumber(allowNeg bool) string {
	if l.atEnd(0) {
		return ""
	}
	l.eatWhitespaces()
	peek := 0
	if allowNeg && l.s[l.c] == '-' {
		peek++
	}
	for !l.atEnd(peek) && isDigit(l.s[l.c+peek]) {
		peek++
	}
	s := l.s[l.c : l.c+peek]
	l.c += peek
	return s
}

func (l *Lexer) nextChars() string {
	if l.atEnd(0) {
		return ""
	}
	l.eatWhitespaces()
	peek := 0
	for !l.atEnd(peek) && !isDigit(l.s[l.c+peek]) && !isWhitespace(l.s[l.c+peek]) {
		peek++
	}
	s := l.s[l.c : l.c+peek]
	l.c += peek
	return s
}

func (l *Lexer) NextStatement() Statement {
	for !l.atEnd(0) {
		num := l.nextNumber(l.negatives)
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

			y, m, d := reorderFormat(n1, n2, n3, l.dateFormat)
			return DateStatement{
				Year:  y,
				Month: m,
				Day:   d,
			}
		} else if nc == ":" { // Parse time
			if strings.HasPrefix(num, "-") {
				num = num[1:]
			}

			hour, _ := strconv.Atoi(num)
			hour %= 24

			num = l.nextNumber(false)
			minute, _ := strconv.Atoi(num)
			minute %= 60

			var second int
			if !l.atEnd(0) && l.s[l.c] == ':' {
				l.c++
				num = l.nextNumber(false)
				second, _ = strconv.Atoi(num)
				second %= 60
			}

			return DaytimeStatement{
				Hour:   hour,
				Minute: minute,
				Second: second,
			}
		} else if nc != "" { // Parse time length
			n, _ := strconv.Atoi(num)
			return TimelengthStatement{
				Value:  n,
				Period: nc,
			}
		}
	}
	return nil
}

func NewLexer(s string, dateFormat dateformat.DateFormat, negatives bool) Lexer {
	return Lexer{
		0, s, dateFormat, negatives,
	}
}
