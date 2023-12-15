package lexer

import (
	"go-interpreter/token"
)

/* Type Def */

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

/* Factory */

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func NewToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: token.ASSIGN, Literal: "="}
}

/* Public Methods */

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			t = token.Token{Type: token.EQ, Literal: literal}
		} else {
			t = token.Token{Type: token.ASSIGN, Literal: "="}
		}
	case '(':
		t = token.Token{Type: token.LPAREN, Literal: "("}
	case ')':
		t = token.Token{Type: token.RPAREN, Literal: ")"}
	case ',':
		t = token.Token{Type: token.COMMA, Literal: ","}
	case '+':
		t = token.Token{Type: token.PLUS, Literal: "+"}
	case '{':
		t = token.Token{Type: token.LBRACE, Literal: "{"}
	case '}':
		t = token.Token{Type: token.RBRACE, Literal: "}"}
	case '-':
		t = token.Token{Type: token.MINUS, Literal: "-"}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			t = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			t = token.Token{Type: token.BANG, Literal: "!"}
		}
	case '*':
		t = token.Token{Type: token.ASTERISK, Literal: "*"}
	case '/':
		t = token.Token{Type: token.FSLASH, Literal: "/"}
	case '\\':
		t = token.Token{Type: token.BSLASH, Literal: "\\"}
	case '<':
		t = token.Token{Type: token.LT, Literal: "<"}
	case '>':
		t = token.Token{Type: token.GT, Literal: ">"}
	case ';':
		t = token.Token{Type: token.SEMICOLON, Literal: ";"}
	case 0:
		t = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
		}
	}
	l.readChar()
	return t
}

/* Private Methods */

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	return l.readMutli(isLetter)
}

func (l *Lexer) readNumber() string {
	return l.readMutli(isDigit)
}

func (l *Lexer) readMutli(validTest func(byte) bool) string {

	position := l.position

	for validTest(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]

}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/* Utils */

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
