package lexer

import "monkey/token"

type Lexer struct {
	input   string
	pos     int  // curr pos in input (curr char)
	readPos int  // curr reading pos in input (curr char + 1)
	ch      byte // curr char under exam
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0 // ascii code to NUL
	} else {
		l.ch = l.input[l.readPos] // set curr char to the next one
	}
	// incr. pos
	l.pos = l.readPos
	l.readPos += 1
}

func newTok(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextTok() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newTok(token.ASSIGN, l.ch)
	case '+':
		tok = newTok(token.PLUS, l.ch)
	case ',':
		tok = newTok(token.COMMA, l.ch)
	case ';':
		tok = newTok(token.SEMICOLON, l.ch)
	case '(':
		tok = newTok(token.LPAREN, l.ch)
	case ')':
		tok = newTok(token.RPAREN, l.ch)
	case '{':
		tok = newTok(token.LBRACE, l.ch)
	case '}':
		tok = newTok(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}
