package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextTok(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expType token.TokenType
		expLit  string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextTok()

		if tok.Type != tt.expType {
			t.Fatalf("%d TokenType wrong, expected %q, got %q", i, tt.expType, tok.Type)
		}

		if tok.Literal != tt.expLit {
			t.Fatalf("%d Token Literal wrong, expected %q, got %q", i, tt.expLit, tok.Literal)
		}
	}
}
