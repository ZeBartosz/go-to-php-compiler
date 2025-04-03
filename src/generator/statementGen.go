package generator

import (
	"strings"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
)

// generateExpressionStmt: Handles expression statements
func generateExpressionStmt(stmt ast.ExpressionStmt) string {
	return generateExpression(stmt.Expression) + ";\n"
}

func generateParams(params []ast.Parameters) string {
	var phpCode strings.Builder

	for i, param := range params {
		phpCode.WriteString(ast.TypeToString(param.Type))
		phpCode.WriteString(" ")
		phpCode.WriteString("$" + param.Name)
		if i < len(params)-1 {
			phpCode.WriteString(", ")
		}
	}

	return phpCode.String()

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

func generateFuncStmt(stmt ast.FuncStmt) string {
	var phpCode strings.Builder

	phpCode.WriteString("public function " + stmt.FuncName + "(" + generateParams(stmt.Params) + ")")
	if stmt.Type != nil {
		phpCode.WriteString(": " + ast.TypeToString(stmt.Type))
	}

	phpCode.WriteString("\n{ \n" + generateBlockStmt(stmt.Block) + "\n }")

	return phpCode.String()
}

func generateReturnStmt(stmt ast.ReturnStmt) string {
	var phpCode strings.Builder

	phpCode.WriteString("return " + generateExpression(stmt.Value) + ";")

	return phpCode.String()
}
