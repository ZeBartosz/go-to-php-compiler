package lexer

import (
	"fmt"
	"regexp"
)

// regexHandler is a function type that processes matched regex patterns.
// It receives a pointer to the lexer and the regex pattern that matched.
type regexHandler func(lex *lexer, regex *regexp.Regexp)

// regexPattern represents a pattern that the lexer will use to tokenize input.ointer to a lexer instance.
// Each pattern consists of a compiled regular expression and a handler function
// that determines what happens when a match is found.
type regexPattern struct {
	regex   *regexp.Regexp // The compiled regular expression used for matching tokens.
	handler regexHandler   // Function to handle the matched token.
}

// lexer is responsible for breaking the source string into tokens using regex patterns.
type lexer struct {
	patterns []regexPattern // List of regex patterns and their associated handlers.
	Tokens   []Token        // List of tokens extracted from the source input.
	source   string         // The input string being tokenized.
	pos      int            // The current position in the source string.
}

// updates the currect positiontokens[p.pos] in the source
func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

// add new token to the array
func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

// checks what is at the current pos
func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

// checks if we are at the end of the source
func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

// checks how many bytes are left till the end
func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	// Iterate while we still have tokens
	for !lex.at_eof() {

		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		// Can extend this print to show location and other stuff
		if !matched {
			panic(fmt.Sprintf("Lexer::Error -> unrecognized token near %s\n", lex.remainder()))
		}
	}

	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
}

// default handling
func defaultHandler(kind TokenKind, value string) regexHandler {
	// pointer to lexer instance
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func createLexer(source string) *lexer {
	// & passing a pointer to a lexer instance
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`\/\/.*`), skipHandler},
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`==`), defaultHandler(EQUALS, "==")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGNMENT, "=")},
			{regexp.MustCompile(`:=`), defaultHandler(ASSIGNMENT, ":=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUALS, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUALS, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")},
			{regexp.MustCompile(`&&`), defaultHandler(AND, "&&")},
			{regexp.MustCompile(`\.\.`), defaultHandler(DOT_DOT, "..")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMI_COLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`--`), defaultHandler(MINUS_MINUS, "--")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(DASH, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
		},
	}
}

// handles numbers, pointer to a lexer instance, a compiled expression
func numberHandler(lex *lexer, regex *regexp.Regexp) {
	// finds first match of number pattern
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}

// handles the stings
func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1 : match[1]-1]

	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceN(len(stringLiteral) + 2)
}

// handles reserved or indentifiers
func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())

	// looks up the match if it exists in the reserved_lu
	if kind, exists := reserved_lu[match]; exists {
		lex.push(NewToken(kind, match))
	} else {
		lex.push(NewToken(IDENTIFIER, match))
	}

	lex.advanceN(len(match))
}

// handles whitespace
func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}
