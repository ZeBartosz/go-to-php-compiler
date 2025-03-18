package ast

// types for ast
type Stmt interface {
	stmt()
}

type Expr interface {
	expr()
}
