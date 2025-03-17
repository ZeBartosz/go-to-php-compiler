package lexer

// Defining an enum-like type
type TokenKind int

const (
	// ioat a special identifier which auto-increments starting from 0
	EOF TokenKind = iota
	NUMBER
	STRING
	IDETIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	ASSIGNMENT // =
	EQUALS     // ==
	NOT
	NOT_EQUALS
)

type token struct {
	// the token type
	kind TokenKind
	// this is the value of the token (eg.+ would have a underlining + value)
	value string
}
