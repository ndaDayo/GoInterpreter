package lexer

import (
    "testing"
    "monkey/token"
)

func TestNextToken(t *testing.T){
    input := `=+(){},;`

    tests := []struct{
        expectedType  token.TokenType
        expectedLiteral string
    }{
        {token.ASSIGN, "="},
        {token.PLUS, "+"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
    }

    l := New(input)

    for i, tt := range tests{
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wring. expected=%q, got=%q",
        i, tt.expectedType, tok.Type)
        }
    }
}
