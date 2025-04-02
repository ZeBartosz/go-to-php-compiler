package generator

import (
	"fmt"
	"strings"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
	"github.com/ZeBartosz/go-to-php-compiler/src/lexer"
)

// GeneratePHP:  The main function to generate PHP code from the AST.
func GeneratePHP(node ast.Stmt) string {
	var phpCode strings.Builder

	phpCode.WriteString("<?php\n\n")

	// Generate code for the main part of the AST.
	phpCode.WriteString(generateStatement(node))

	return phpCode.String()
}

// generateStatement: Handles different statement types
func generateStatement(stmt ast.Stmt) string {
	switch n := stmt.(type) {
	case ast.BlockStmt: // Corrected type name
		return generateBlockStmt(n)
	case ast.ExpressionStmt:
		return generateExpressionStmt(n)
	case ast.VarDeclStmt:
		return generateVarDeclStmt(n)
	default:
		return fmt.Sprintf("// Unsupported statement type: %T\n", stmt)
	}
}

// generateExpression:  Handles different expression types.
func generateExpression(expr ast.Expr) string {
	switch n := expr.(type) {
	case ast.NumberExpr:
		return fmt.Sprintf("%g", n.Value)
	case ast.StringExpr:
		return fmt.Sprintf("%q", n.Value)
	case ast.SymbolExpr:
		return "$" + n.Value
	case ast.BinaryExpr:
		return generateBinaryExpr(n)
	case ast.PrefixExpr:
		return generatePrefixExpr(n)
	case ast.AssignmentExpr:
		return generateAssignmentExpr(n)
	default:
		return fmt.Sprintf("// Unsupported expression type: %T\n", expr)
	}
}

// translateOperator: Translates Go operators to PHP operators.
func translateOperator(token lexer.Token) string {
	switch token.Kind {
	case lexer.PLUS:
		return "+"
	case lexer.DASH:
		return "-"
	case lexer.STAR:
		return "*"
	case lexer.SLASH:
		return "/"
	case lexer.EQUALS:
		return "="
	case lexer.GREATER:
		return ">"
	default:
		return "// Unsupported operator"
	}
}
