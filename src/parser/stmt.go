package parser

import (
	"github.com/ZeBartosz/PrattParsing/src/ast"
	"github.com/ZeBartosz/PrattParsing/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {

	// check if the stmt of the token exists
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	// if it does return stmt function
	if exists {
		return stmt_fn(p)
	}

	// parse the token
	expression := parse_expr(p, defalt_bp)
	// checks if its a semicolon
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}

func parse_var_decl_stmt(p *parser) ast.Stmt {
	// checks if the token is a const
	isConst := p.advance().Kind == lexer.CONST
	// checks for the value
	varName := p.expectError(lexer.IDENTIFIER, "Inside the variable declaration expected to find value").Value
	// expect the current token to be an assignment
	p.expect(lexer.ASSIGNMENT)

	assignedValue := parse_expr(p, assignment)
	p.expect(lexer.SEMI_COLON)

	return ast.VarDeclStmt{
		IsConstant:    isConst,
		VariableName:  varName,
		AssignedValue: assignedValue,
	}
}
