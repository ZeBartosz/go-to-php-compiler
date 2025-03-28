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

func Parse(tokens []lexer.Token) (ast.Stmt, error) {

	// Create the body
	p := createParser(tokens)

	// deals with the package
	if p.currentTokenKind() == lexer.PACKAGE {
		p.advance()
		p.advance()

		// Debug
		// p.expect(lexer.IDENTIFIER).Value
		// fmt.Printf("Parsed package: %s\n", packageName)
	}

	// instancate a body
	Body := make([]ast.Stmt, 0)

	// iterate till we get to the end of the file
	for p.hasToken() {
		stmt, err := parse_stmt(p)
		if err != nil {
			return nil, err
		}
		Body = append(Body, stmt)
	}

	return ast.BlockStmt{
		Body: Body,
	}, nil
}

// helper methods
// returns current token
func (p *parser) currentToken() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{Kind: lexer.EOF}
	}
	return p.tokens[p.pos]
}

// we get the current kind of the token
func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

// advance to the next token, returns current token
func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	if p.pos < len(p.tokens) {
		p.pos++
	}
	return tk
}

// Checks if there are more token to parse through
func (p *parser) hasToken() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

// checks if the expected tokenKind is the same a the parse kind
func (p *parser) expectError(expectedKind lexer.TokenKind, errStr string) (lexer.Token, error) {
	token := p.currentToken()
	kind := token.Kind

	if kind != expectedKind {
		err := fmt.Errorf("expected %s but received %s instead", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))
		if errStr != "" {
			err = fmt.Errorf("%s: %w", errStr, err)
		}
		return token, err
	}

	return p.advance(), nil
}

// checks if the token is the one we expect
func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	tok, err := p.expectError(expectedKind, "")
	if err != nil {
		panic(err)
	}
	return tok
}
