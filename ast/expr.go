package ast

import "MrTeeny/token"

// Node es solo una interface sin nada.
type Node interface{}

type VisitorExpr interface {
	VisitLiteral(e *Literal) Node
	VisitUnary(e *PrefixExpr) Node
	VisitBinary(e *InfixExpr) Node
}

type Expr interface {
	// OJO: ¿por qué una copia del visitor?
	Accept(v VisitorExpr) Node
}

// Literal: foo, true, 12, "bar"
type Literal struct {
	Token token.Token
	value interface{}
}

func (e *Literal) Accept(v VisitorExpr) Node {
	return v.VisitLiteral(e)
}

type PrefixExpr struct {
	Token    token.Token
	Operator string
	Right    Expr
}

func (e *PrefixExpr) Accept(v VisitorExpr) Node {
	return v.VisitUnary(e)
}

type InfixExpr struct {
	Token    token.Token
	Left     Expr
	Operator string
	Right    Expr
}

func (e *InfixExpr) Accept(v VisitorExpr) Node {
	return v.VisitBinary(e)
}
