package timeless

type lexer struct {
	c   int
	s   string
	len int
}

func (l *lexer) atEnd(peek int) bool {
	return l.c+peek == l.len
}

func (l *lexer) eatWhitespaces() {
	for isWhitespace(l.s[l.c]) {
		l.c++
	}
}

func (l *lexer) nextNumber(allowNeg bool) string {
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

func (l *lexer) nextChars() string {
	l.eatWhitespaces()
	peek := 0
	for !l.atEnd(peek) && !isDigit(l.s[l.c+peek]) && !isWhitespace(l.s[l.c+peek]) {
		peek++
	}
	s := l.s[l.c : l.c+peek]
	l.c += peek
	return s
}

func newLexer(s string) lexer {
	return lexer{
		0, s, len(s),
	}
}
