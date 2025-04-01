package parser

import (
	"fmt"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
	"github.com/ZeBartosz/go-to-php-compiler/src/lexer"
)

func parse_stmt(p *parser) (ast.Stmt, error) {
	// check if the stmt of the token exists
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	// if it does return stmt function
	if exists {
		return stmt_fn(p)
	}

	// parse the token
	expression, err := parse_expr_stmt(p)
	if err != nil {
		return nil, err
	}

	return expression, nil
}

func parse_expr_stmt(p *parser) (ast.Stmt, error) {
	expression := parse_expr(p, defalt_bp)
	// checks if its a eof
	_, err := p.expectError(lexer.EOF, "Expected EOF at end of expression")
	if err != nil {
		return nil, err
	}

	return ast.ExpressionStmt{
		Expression: expression,
	}, nil
}

func parse_var_decl_stmt(p *parser) (ast.Stmt, error) {
	var explicitType ast.Type
	var assignedValue ast.Expr

	// checks if the token is a const
	isConst := p.advance().Kind == lexer.CONST
	// checks for the value
	ident, err := p.expectError(lexer.IDENTIFIER, "Inside the variable declaration expected to find value")
	if err != nil {
		return nil, err
	}

	varName := ident.Value

	if p.currentTokenKind() == lexer.COLON {
		p.advance()
		explicitType = parse_type(p, defalt_bp)
	}

	if p.currentTokenKind() != lexer.SEMI_COLON {
		// expect the current token to be an assignment
		_, err = p.expectError(lexer.ASSIGNMENT, "Expected assignment operator (=)")
		if err != nil {
			return nil, err
		}
		assignedValue = parse_expr(p, assignment)
	} else if explicitType == nil {
		return nil, fmt.Errorf("missing either right-hand side of var declaration or explicit type")
	}

	_, err = p.expectError(lexer.SEMI_COLON, "Expected semicolon at the end of variable declaration")
	if err != nil {
		return nil, err
	}

	if isConst && assignedValue == nil {
		return nil, fmt.Errorf("cannot define constant without providing value")
	}

	return ast.VarDeclStmt{
		ExplicitType:  explicitType,
		IsConstant:    isConst,
		VariableName:  varName,
		AssignedValue: assignedValue,
	}, nil
}

// Function to parse import statements
func parse_import_stmt(p *parser) (ast.Stmt, error) {
	p.advance()

	// Expect the next token to be an identifier (package name)
	if p.currentTokenKind() != lexer.STRING {
		return nil, fmt.Errorf("expected package name after import, current package name: %s and kind: %s", p.currentToken().Value, lexer.TokenKindString(p.currentTokenKind()))
	}

	packageName := p.advance().Value

	return ast.ImportStmt{
		PackageName: packageName,
	}, nil
}
