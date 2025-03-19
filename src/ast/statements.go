package ast

// { ... } anything inside of the block of brackets
type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) stmt() {}

type VarDeclStmt struct {
	VariableName  string
	IsConstant    bool
	AssignedValue Expr
}

func (n VarDeclStmt) stmt() {}
