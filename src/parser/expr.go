package parser

import (
	"fmt"
	"strconv"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
	"github.com/ZeBartosz/go-to-php-compiler/src/lexer"
)

func parse_expr(p *parser, bp binding_power) ast.Expr {
	// gets the current token kind
	tokenKind := p.currentTokenKind()

	// checks nud look up for the token kind
	nud_fn, exists := nud_lu[tokenKind]

	// if it doesnt exists through an error
	if !exists {
		panic(fmt.Sprintf("NUD HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)
	// checks if the current tokenKind has higher binding power
	for bp_lu[p.currentTokenKind()] > bp {
		// next tokenKind
		tokenKind = p.currentTokenKind()
		// checks if it exists in the look up
		led_fn, exists := led_lu[tokenKind]

		// if it doesnt provide an error
		if !exists {
			panic(fmt.Sprintf("LED HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp_lu[p.currentTokenKind()])
	}

	return left
}

func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		if p.containsFunc(p.currentToken().Value) {
			return parse_func_call_expr(p)
		} else if p.containsImport(p.currentToken().Value) {
			return parse_import_expr(p)
		}
		return ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Cannot create primary_expr from %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

func parse_import_expr(p *parser) ast.Expr {
	p.advance()
	p.expect(lexer.DOT)

	importValue := p.advance().Value

	p.expect(lexer.OPEN_PAREN)
	var params []ast.ImportParams
	for p.currentTokenKind() != lexer.CLOSE_PAREN {
		isIdentifier := false
		if p.currentToken().Kind == lexer.IDENTIFIER {
			isIdentifier = true
		}

		paramValue := p.advance().Value

		params = append(params, ast.ImportParams{
			Identifier: isIdentifier,
			Value:      paramValue,
		})

		if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		}
	}

	p.expect(lexer.CLOSE_PAREN)

	return ast.ImportCallExpr{
		Value:   importValue,
		Pararms: params,
	}
}

func parse_func_call_expr(p *parser) ast.Expr {
	funcName := p.advance().Value

	p.expect(lexer.OPEN_PAREN)

	var params []string
	for p.currentTokenKind() != lexer.CLOSE_PAREN {
		paramName := p.expect(lexer.IDENTIFIER).Value

		params = append(params, paramName)

		if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		}
	}

	p.expect(lexer.CLOSE_PAREN)

	return ast.FuncCallExpr{
		Value:   funcName,
		Pararms: params,
	}
}

func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()
	right := parse_expr(p, defalt_bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

func parse_prefix_expr(p *parser) ast.Expr {
	operatorToken := p.advance()
	rhs := parse_expr(p, defalt_bp)

	return ast.PrefixExpr{
		Operator:  operatorToken,
		RightExpr: rhs,
	}

}

func parse_grouping_expr(p *parser) ast.Expr {
	p.advance()
	expr := parse_expr(p, defalt_bp)
	p.expect(lexer.CLOSE_PAREN)

	return expr
}

func parse_assignment_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()
	rhs := parse_expr(p, bp)

	return ast.AssignmentExpr{
		Assigne:  left,
		Operator: operatorToken,
		Value:    rhs,
	}

}
