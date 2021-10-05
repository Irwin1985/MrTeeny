package token

import "fmt"

// Type es el tipo de todos los tokens
// ejemplo:
//	INT es un Type para los entero
// 	STRING es un Type para las cadena

type Type int

// Lista de Tokens
const (
	ILLEGAL Type = iota
	EOF

	// Identificadores y literales
	IDENT
	INT
	STRING

	// Operadores aritméticos
	ASSIGN   // '='
	PLUS     // '+'
	PLUS_EQ  // '+='
	MINUS    // '-'
	MINUS_EQ // '-='
	MUL      // '*'
	MUL_EQ   // '*='
	DIV      // '/'
	DIV_EQ   // '/='

	// Operadores relacionales
	LT  // '<'
	LEQ // '<='
	GT  // '>'
	GEQ // '>='
	EQ  // '=='
	NEQ // '!='
	NOT // '!'

	// Caracteres especiales
	COMMA
	SEMICOLON
	COLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	// Palabras reservadas
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

// Constantes literales para los tipos de token
var tokenNames = []string{
	"ILLEGAL",
	"EOF",

	// Identificadores y literales
	"IDENT",
	"INT",
	"STRING",

	// Operadores aritméticos
	"ASSIGN",   // '='
	"PLUS",     // '+'
	"PLUS_EQ",  // '+='
	"MINUS",    // '-'
	"MINUS_EQ", // '-='
	"MUL",      // '*'
	"MUL_EQ",   // '*='
	"DIV",      // '/'
	"DIV_EQ",   // '/='

	// Operadores relacionales
	"LT",  // '<'
	"LEQ", // '<='
	"GT",  // '>'
	"GEQ", // '>='
	"EQ",  // '=='
	"NEQ", // '!='
	"NOT", // '!'

	// Caracteres especiales
	"COMMA",
	"SEMICOLON",
	"COLON",

	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",
	"LBRACKET",
	"RBRACKET",

	// Palabras reservadas
	"FUNCTION",
	"LET",
	"TRUE",
	"FALSE",
	"IF",
	"ELSE",
	"RETURN",
}

// Token es la estructura base para el parser. Es creada en el Lexer y se usa
// durante el proceso de interpretación del lenguaje.
type Token struct {
	Type   Type
	Lexeme string
}

func (t Token) ToString() string {
	return fmt.Sprintf("<%s, '%s'>", tokenNames[t.Type], t.Lexeme)
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
