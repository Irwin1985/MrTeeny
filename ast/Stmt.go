package ast

import "MrTeeny/token"

type VisitorStmt interface {
	VisitLetStmt(s *LetStmt) Node
	VisitReturnStmt(s *ReturnStmt) Node
}

type Stmt interface {
	Accept(v VisitorStmt) Node
}

type LetStmt struct {
	Token token.Token
	Name  Literal
	Value Expr
}

func (s *LetStmt) Accept(v VisitorStmt) Node {
	return v.VisitLetStmt(s)
}

type ReturnStmt struct {
	Token token.Token
	Value Expr
}

func (s *ReturnStmt) Accept(v VisitorStmt) Node {
	return v.VisitReturnStmt(s)
}
