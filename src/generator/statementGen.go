package generator

import (
	"strings"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
)

// generateExpressionStmt: Handles expression statements
func generateExpressionStmt(stmt ast.ExpressionStmt, gen *Generator) {
	gen.writeln(generateExpression(stmt.Expression) + ";\n")
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
func generateVarDeclStmt(stmt ast.VarDeclStmt, gen *Generator) {
	gen.writeln("$" + stmt.VariableName + " = " + generateExpression(stmt.AssignedValue) + ";\n")
}

// generateBlockStmt: Handles block statements
func generateBlockStmt(block ast.BlockStmt, gen *Generator) {
	for _, stmt := range block.Body {
		generateStatement(stmt, gen)
	}
}

func generateFuncStmt(stmt ast.FuncStmt, gen *Generator) {

	gen.increaseIndent()

	if stmt.Type != nil {
		gen.writeln("public function " + stmt.FuncName + "(" + generateParams(stmt.Params) + ")" + ": " + ast.TypeToString(stmt.Type))
	} else {
		gen.writeln("public function " + stmt.FuncName + "(" + generateParams(stmt.Params) + ")")
	}

	gen.writeln("{")
	gen.increaseIndent()
	generateBlockStmt(stmt.Block, gen)
	gen.decreaseIndent()
	gen.writeln("}")

	gen.decreaseIndent()
}

func generateReturnStmt(stmt ast.ReturnStmt, gen *Generator) {

	gen.writeln("return " + generateExpression(stmt.Value) + ";")

}
