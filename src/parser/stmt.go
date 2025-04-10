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
	// _, err := p.expectError(lexer.EOF, "Expected EOF at end of expression")
	// if err != nil {
	// 	return nil, err
	// }

	return ast.ExpressionStmt{
		Expression: expression,
	}, nil
}

func parse_return_stmt(p *parser) (ast.Stmt, error) {
	p.advance()

	returnValue := parse_expr(p, defalt_bp)

	return ast.ReturnStmt{
		Value: returnValue,
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

	if p.currentTokenKind() != lexer.ASSIGNMENT {
		explicitType = parse_type(p, defalt_bp)
	}

	if p.currentTokenKind() == lexer.ASSIGNMENT {
		// expect the current token to be an assignment
		_, err = p.expectError(lexer.ASSIGNMENT, "Expected assignment operator (=)")
		if err != nil {
			return nil, err
		}

		assignedValue = parse_expr(p, assignment)
	}

	if explicitType == nil {
		return nil, fmt.Errorf("missing either right-hand side of var declaration or explicit type")
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
	p.addImport(packageName)

	return ast.ImportStmt{
		PackageName: packageName,
	}, nil
}

func parse_func_stmt(p *parser) (ast.Stmt, error) {
	p.advance()
	funcName := p.expect(lexer.IDENTIFIER).Value
	p.expect(lexer.OPEN_PAREN)

	var params []ast.Parameters
	for p.currentTokenKind() != lexer.CLOSE_PAREN {
		paramName := p.expect(lexer.IDENTIFIER).Value
		paramType := parse_type(p, defalt_bp)

		params = append(params, ast.Parameters{Name: paramName, Type: paramType})

		if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		}
	}

	p.expect(lexer.CLOSE_PAREN)

	var funcType ast.Type

	if p.currentTokenKind() != lexer.OPEN_CURLY {
		funcType = parse_type(p, defalt_bp)
	}

	p.expect(lexer.OPEN_CURLY)

	body := make([]ast.Stmt, 0)

	for p.currentTokenKind() != lexer.CLOSE_CURLY {
		stmt, err := parse_stmt(p)
		if err != nil {
			return nil, err
		}
		body = append(body, stmt)
	}

	p.expect(lexer.CLOSE_CURLY)

	blockStmt := ast.BlockStmt{
		Body: body,
	}

	p.addFunction(funcName, params)

	return ast.FuncStmt{
		FuncName: funcName,
		Params:   params,
		Type:     funcType,
		Block:    blockStmt,
	}, nil
}
