package ast

import "github.com/ZeBartosz/PrattParsing/src/lexer"

// LITERAL EXPRESSIONS
type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {}

// COMPLEX  EXPRESSIONS

// 10 + 5	,10 is the left expr, + is the Operator, 5 is the right Expr, right can hold another BinaryExpr
type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}
