package codegen

import (
	"fmt"
	"strings"

	"github.com/ZeBartosz/PrattParsing/src/ast"
	"github.com/ZeBartosz/PrattParsing/src/lexer"
)

// GeneratePHP:  The main function to generate PHP code from the AST.
func GeneratePHP(node ast.Stmt) string {
	var phpCode strings.Builder

	phpCode.WriteString("<?php\n")

	// Generate code for the main part of the AST.
	phpCode.WriteString(generateStatement(node))

	phpCode.WriteString("\n?>")

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

// generateExpressionStmt: Handles expression statements
func generateExpressionStmt(stmt ast.ExpressionStmt) string {
	return generateExpression(stmt.Expression) + ";\n"
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

// generateBinaryExpr: Handles binary expressions.
func generateBinaryExpr(expr ast.BinaryExpr) string {
	left := generateExpression(expr.Left)
	right := generateExpression(expr.Right)
	operator := translateOperator(expr.Operator)
	return fmt.Sprintf("(%s %s %s)", left, operator, right)
}

// generatePrefixExpr: Handles prefix expressions.
func generatePrefixExpr(expr ast.PrefixExpr) string {
	operator := translateOperator(expr.Operator)
	right := generateExpression(expr.RightExpr)
	return fmt.Sprintf("%s%s", operator, right)
}

// generateAssignmentExpr: Handles assignment expressions.
func generateAssignmentExpr(expr ast.AssignmentExpr) string {
	assignee := generateExpression(expr.Assigne)
	operator := translateOperator(expr.Operator)
	value := generateExpression(expr.Value)
	return fmt.Sprintf("%s %s %s", assignee, operator, value)
}

// generateVarDeclStmt: Handles variable declaration statements.
func generateVarDeclStmt(stmt ast.VarDeclStmt) string {
	var phpCode strings.Builder

	phpCode.WriteString("$" + stmt.VariableName + " = " + generateExpression(stmt.AssignedValue) + ";\n")

	return phpCode.String()
}

// generateBlockStmt: Handles block statements
func generateBlockStmt(block ast.BlockStmt) string {
	var phpCode strings.Builder
	for _, stmt := range block.Body {
		phpCode.WriteString(generateStatement(stmt))
	}
	return phpCode.String()
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
	default:
		return "// Unsupported operator"
	}
}
