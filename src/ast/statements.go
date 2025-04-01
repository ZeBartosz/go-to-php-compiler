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
	ExplicitType  Type
}

func (n VarDeclStmt) stmt() {}

type ImportStmt struct {
	PackageName string
}

func (n ImportStmt) stmt() {}

type Parameters struct {
	Name string
	Type Type
}

type FuncStmt struct {
	FuncName string
	Params   []Parameters
	Type     Type
	Block    BlockStmt
}

func (n FuncStmt) stmt() {}

type ReturnStmt struct {
	Value Expr
}

func (n ReturnStmt) stmt() {}
