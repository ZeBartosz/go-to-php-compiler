package generator

import (
	"fmt"
	"strings"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
)

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

func generateFuncCallExpr(expr ast.FuncCallExpr) string {
	params := make([]string, len(expr.Pararms))

	for i, param := range expr.Pararms {
		params[i] = "$" + param
	}

	return fmt.Sprintf("$this->%s(%s)", expr.Value, strings.Join(params, ", "))
}
