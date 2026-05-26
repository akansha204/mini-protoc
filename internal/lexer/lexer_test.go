package lexer

import (
	"testing"

	"github.com/akansha204/mini-protoc/internal/token"
)

func TestNextToken(t *testing.T) {
	input := `
        syntax = "proto3";
        package calculator.math;

		// request message -> A comment 

        message AddRequest {
	        int32 a = 1;
	        int64 b = 2;
	        string user_name = 3;
	        bool active = 4;
	        float x = 5;
	        double y = 6;
        }

        message AddResponse {
	        int32 result = 1;
        }

        service CalculatorService {
	        rpc Add(AddRequest) returns (AddResponse);
	        rpc Multiply(AddRequest) returns (AddResponse);
        } `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SYNTAX, "syntax"},
		{token.ASSIGN, "="},
		{token.STRING, "proto3"},
		{token.SEMICOLON, ";"},

		{token.PACKAGE, "package"},
		{token.IDENT, "calculator"},
		{token.DOT, "."},
		{token.IDENT, "math"},
		{token.SEMICOLON, ";"},

		{token.MESSAGE, "message"},
		{token.IDENT, "AddRequest"},
		{token.LBRACE, "{"},

		{token.IDENT, "int32"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.NUMBER, "1"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "int64"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.NUMBER, "2"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "string"},
		{token.IDENT, "user_name"},
		{token.ASSIGN, "="},
		{token.NUMBER, "3"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "bool"},
		{token.IDENT, "active"},
		{token.ASSIGN, "="},
		{token.NUMBER, "4"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "float"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "double"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.NUMBER, "6"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},

		{token.MESSAGE, "message"},
		{token.IDENT, "AddResponse"},
		{token.LBRACE, "{"},

		{token.IDENT, "int32"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.NUMBER, "1"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},

		{token.SERVICE, "service"},
		{token.IDENT, "CalculatorService"},
		{token.LBRACE, "{"},

		{token.RPC, "rpc"},
		{token.IDENT, "Add"},
		{token.LPAREN, "("},
		{token.IDENT, "AddRequest"},
		{token.RPAREN, ")"},
		{token.RETURNS, "returns"},
		{token.LPAREN, "("},
		{token.IDENT, "AddResponse"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.RPC, "rpc"},
		{token.IDENT, "Multiply"},
		{token.LPAREN, "("},
		{token.IDENT, "AddRequest"},
		{token.RPAREN, ")"},
		{token.RETURNS, "returns"},
		{token.LPAREN, "("},
		{token.IDENT, "AddResponse"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tok.Type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tok.Literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
