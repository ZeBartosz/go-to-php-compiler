package parser

import (
	"fmt"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
	"github.com/ZeBartosz/go-to-php-compiler/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

// create a parser instance
func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	createTokenTypeLookups()

	return &parser{
		tokens: tokens, pos: 0,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	// instancate a body
	Body := make([]ast.Stmt, 0)
	// Create the body
	p := createParser(tokens)

	// iterate till we get to the end of the file
	for p.hasToken() {
		Body = append(Body, parse_stmt(p))
	}

	return ast.BlockStmt{
		Body: Body,
	}
}

// helper methods
// returns current token
func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

// we get the current kind of the token
func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

// advance to the next token, returns current token
func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

// Checks if there are more token to parse through
func (p *parser) hasToken() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

// checks if the expected tokenKind is the same a the parse kind
func (p *parser) expectError(expectedKind lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind

	if kind != expectedKind {
		if err == nil {
			err = fmt.Sprintf("Expected %s but received %s instead\n", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))
		}

		panic(err)
	}

	return p.advance()
}

// checks if the token is the one we expect1
func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}
