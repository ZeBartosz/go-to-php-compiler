package generator

import (
	"fmt"
	"strings"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
	"github.com/ZeBartosz/go-to-php-compiler/src/lexer"
)

type Generator struct {
	indentLevel int
	strBuilder  strings.Builder
}

// GeneratePHP:  The main function to generate PHP code from the AST.
func GeneratePHP(node ast.Stmt) string {
	gen := &Generator{
		indentLevel: 0,
	}

	gen.writeln("<?php\n")
	gen.writeln("class Main\n" + "{\n")

	// Generate code for the main part of the AST.
	generateStatement(node, gen)

	gen.writeln("}")

	return gen.strBuilder.String()
}

// generateStatement: Handles different statement types
func generateStatement(stmt ast.Stmt, gen *Generator) {
	switch n := stmt.(type) {
	case ast.BlockStmt:
		generateBlockStmt(n, gen)
	case ast.ExpressionStmt:
		generateExpressionStmt(n, gen)
	case ast.VarDeclStmt:
		generateVarDeclStmt(n, gen)
	case ast.FuncStmt:
		generateFuncStmt(n, gen)
	case ast.ReturnStmt:
		generateReturnStmt(n, gen)
	default:
		gen.writeln(fmt.Sprintf("// Unsupported statement type: %T\n", stmt))
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
	case ast.FuncCallExpr:
		return generateFuncCallExpr(n)
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
	case lexer.ASSIGNMENT:
		return "="
	default:
		return "// Unsupported operator"
	}
}
