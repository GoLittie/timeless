package lexer

type StatementType = int

const (
	TimelengthStatementType StatementType = iota
	DaytimeStatementType
	DateStatementType
	YearStatementType
	MonthStatementType
	DayStatementType
)

type Statement interface {
	Type() StatementType
}

type TimelengthStatement struct {
	Statement
	Value  int
	Period string
}

func (s TimelengthStatement) Type() StatementType {
	return TimelengthStatementType
}

type DaytimeStatement struct {
	Statement
	Hour   int
	Minute int
	Second int
}

func (s DaytimeStatement) Type() StatementType {
	return DaytimeStatementType
}

type DateStatement struct {
	Statement
	Year  int
	Month int
	Day   int
}

func (s DateStatement) Type() StatementType {
	return DateStatementType
}
