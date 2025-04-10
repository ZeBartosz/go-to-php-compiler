package parser

import (
	"fmt"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
	"github.com/ZeBartosz/go-to-php-compiler/src/lexer"
)

type parser struct {
	tokens     []lexer.Token
	pos        int
	funcList   []ast.FuncInfo
	ImportList []ast.ImportStmt
}

// create a parser instance
func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	createTokenTypeLookups()

	return &parser{
		tokens:   tokens,
		pos:      0,
		funcList: []ast.FuncInfo{},
	}
}

func Parse(tokens []lexer.Token) (ast.Stmt, error) {
	// Create the parser instance
	p := createParser(tokens)

	// Handle the package declaration
	if p.currentTokenKind() == lexer.PACKAGE {
		p.advance()
		if p.currentTokenKind() != lexer.IDENTIFIER {
			return nil, fmt.Errorf("expected package name after 'package'")
		}

		p.advance()
		// packageName := p.advance().Value
		// fmt.Printf("Parsed package: %s\n", packageName)
	}

	Body := make([]ast.Stmt, 0)

	// Iterate until we reach the end of the file
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
