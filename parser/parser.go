package parser

import "MrTeeny/ast"
import "MrTeeny/token"

// Orden de precedencia
const (
	_ int = iota
	LOWEST
	EQUALITY
	COMPARISON
	TERM
	FACTOR
	PREFIX
	CALL
	INDEX
)

// Tipos de funciones para relacionar con tokens
type PrefixParsingFn func() ast.Expr
type InfixParsingFn func(ast.Expr) ast.Expr

// Diccionario de tokens con precedencias
var precedences = map[token.Type]int{
	token.EQ:       EQUALITY,
	token.NOT:      EQUALITY,
	token.LT:       COMPARISON,
	token.LEQ:      COMPARISON,
	token.GT:       COMPARISON,
	token.GEQ:      COMPARISON,
	token.PLUS:     TERM,
	token.PLUS_EQ:  TERM,
	token.MINUS:    TERM,
	token.MINUS_EQ: TERM,
	token.MUL:      FACTOR,
	token.MUL_EQ:   FACTOR,
	token.DIV:      FACTOR,
	token.DIV_EQ:   FACTOR,
	token.LPAREN:   PREFIX,
	token.LBRACKET: INDEX,
}

type Parser struct {
}
