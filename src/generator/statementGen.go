package generator

import (
	"strings"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
)

// generateExpressionStmt: Handles expression statements
func generateExpressionStmt(stmt ast.ExpressionStmt) string {
	return generateExpression(stmt.Expression) + ";\n"
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
