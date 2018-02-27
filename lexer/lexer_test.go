package lexer

import (
	"testing"
	"github.com/geertvl/taiki/token"
)

func TestNextToken(t *testing.T) {
	input := `=(){},`

	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if (tok.Type != tt.expectedType) {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if (tok.Literal != tt.expectedLiteral) {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestFreeDiagram(t *testing.T) {
}

func TestSequenceDiagram(t *testing.T) {	
	input := `
		diagram {
			actor act;
			object obj;

			act -> obj;
			obj => act;
		}
	`
	
	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral string
	}{
		{token.DIAGRAM, "diagram"},
		{token.LBRACE, "{"},		
		{token.IDENT, "actor"},
		{token.IDENT, "act"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "object"},
		{token.IDENT, "obj"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "act"},
		{token.SEND, "->"},
		{token.IDENT, "obj"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "obj"},
		{token.RETURN, "=>"},
		{token.IDENT, "act"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if (tok.Type != tt.expectedType) {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if (tok.Literal != tt.expectedLiteral) {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestInclude(t *testing.T) {
	input := `
		#include "other_diagram";

		diagram {

		}
	`
	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral string
	}{
		{ token.INCLUDE, "include" },
		{ token.STRING, "other_diagram"  },
		{ token.SEMICOLON, ";" },
		{ token.DIAGRAM, "diagram" },
		{ token.LBRACE, "{" },
		{ token.RBRACE, "}" },
		{ token.EOF, "" },	
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if (tok.Type != tt.expectedType) {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if (tok.Literal != tt.expectedLiteral) {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}