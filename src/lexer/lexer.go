package lexer

import (
	"fmt"
	"regexp"
)

// regexHandler is a function type that processes matched regex patterns.
// It receives a pointer to the lexer and the regex pattern that matched.
type regexHandler func(lex *lexer, regex *regexp.Regexp)

// regexPattern represents a pattern that the lexer will use to tokenize input.
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

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

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

	return lex.Tokens
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
		},
	}
}
