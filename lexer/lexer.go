package lexer

import (
	"MrTeeny/token"
	"unicode"
)

type Lexer struct {
	input []rune // un slice de caracteres rune
	pos   int    // la posición del caracter actual
	next  int    // la posición del caracter siguiente
	ch    rune
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.advance()
	return l
}

/*
* FUNCIONES HELPER
 */
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
}

func isIdentifier(ch rune) bool {
	return isLetter(ch) || isDigit(ch) || ch == '_'
}

func isWhitespace(ch rune) bool {
	return ch == rune(' ') || ch == rune('\t') || ch == rune('\n') || ch == rune('\r')
}

// avanza la posición al siguiente caracter
// si ha llegado al final de la cadena retorna 0
func (l *Lexer) advance() {
	if l.next >= len(l.input) {
		l.ch = rune(0)
	} else {
		l.ch = l.input[l.next]
	}
	l.pos = l.next
	l.next += 1
}

// salta los espacios en blanco
func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.advance()
	}
}

// obtiene un número
func (l *Lexer) readNumber() string {
	start := l.pos
	for isDigit(l.ch) {
		l.advance()
	}
	return string(l.input[start:l.pos])
}

// obtiene un string
func (l *Lexer) readString() string {
	delim := l.ch
	start := l.pos
	l.advance() // avanza el primer delimitador
	for {
		if l.ch == delim || l.ch == 0 {
			break
		}
	}
	return string(l.input[start:l.pos])
}

// lee un identificador
func (l *Lexer) readIdentifier() string {
	start := l.pos
	for isIdentifier(l.ch) {
		l.advance()
	}
	return string(l.input[start:l.pos])
}

func (l *Lexer) peek() rune {
	if l.next < len(l.input) {
		return l.input[l.next]
	}
	return rune(0)
}

func newToken(tok token.Type, lex string) token.Token {
	return token.Token{Type: tok, Lexeme: lex}
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	var tok token.Token
	switch l.ch {
	case '=':
		if l.peek() == rune('=') {
			l.advance()
			tok = newToken(token.EQ, "==")
		} else {
			tok = newToken(token.ASSIGN, "=")
		}
	case '+':
		if l.peek() == rune('=') {
			l.advance()
			tok = newToken(token.PLUS_EQ, "+=")
		} else {
			tok = newToken(token.PLUS, "+")
		}
	case '-':
		if l.peek() == rune('=') {
			l.advance()
			tok = newToken(token.MINUS_EQ, "-=")
		} else {
			tok = newToken(token.MINUS, "-")
		}
	case '*':
		if l.peek() == rune('=') {
			l.advance()
			tok = newToken(token.MUL_EQ, "*=")
		} else {
			tok = newToken(token.MUL, "*")
		}
	case '/':
		if l.peek() == rune('=') {
			l.advance()
			tok = newToken(token.DIV_EQ, "/=")
		} else {
			tok = newToken(token.DIV, "/")
		}
	case '"':
		tok = newToken(token.STRING, l.readString())
	case ',':
		tok = newToken(token.COMMA, ",")
	case ':':
		tok = newToken(token.COLON, ":")
	case ';':
		tok = newToken(token.SEMICOLON, ";")
	case '(':
		tok = newToken(token.LPAREN, "(")
	case ')':
		tok = newToken(token.RPAREN, ")")
	case '{':
		tok = newToken(token.LBRACE, "{")
	case '}':
		tok = newToken(token.RBRACE, "}")
	case '[':
		tok = newToken(token.LBRACKET, "[")
	case ']':
		tok = newToken(token.RBRACKET, "]")
	case 0:
		tok = newToken(token.EOF, "")
	default:
		// los identificadores y enteros deben retornarse
		// antes de llegar al final del switch porque ellos
		// leen hasta identificar el token. Ejemplo:
		// 123+45 => el primer token es 123 y si dejamos
		// el retorno al final entonces l.advance() se comería
		// el token '+' lo cual está mal. Lo mismo para ident.
		if isLetter(l.ch) {
			tok.Lexeme = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Lexeme)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Lexeme = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch))
		}
	}
	l.advance()
	return tok
}
