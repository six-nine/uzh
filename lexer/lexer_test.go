package lexer

import (
	"fmt"
	"testing"

	"github.com/six-nine/uzh/token"
)

func TestTokenizing(t *testing.T) {
	input := "int a = 1 + b;"
	expected := []token.Token{
		{token.INT, "int"},
		{token.IDENTIFIER, "a"},
		{token.ASSIGN, "="},
		{token.INT_LITERAL, "1"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "b"},
		{token.SEMICOLON, ";"},
	}

	l := New()

	tokens := l.Tokenize(input)

	for i, tok := range tokens {
		fmt.Println(tok.Type, tok.Literal)
		if tok.Type != expected[i].Type {
			t.Fatalf("test failed at token [%d]. Types are different. Expected %q, found %q", i, expected[i].Type, tok.Type)
		}
		if tok.Literal != expected[i].Literal {
			t.Fatalf("test failed at token [%d]. Literals are different. Expected %q, found %q", i, expected[i].Literal, tok.Literal)
		}
	}

}
