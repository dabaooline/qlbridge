package qlparser

import (
	//u "github.com/araddon/gou"
	"github.com/bmizerany/assert"
	"testing"
)

func tokenexpr(lexString string, runLex StateFn) Token {
	l := NewLexer(lexString, ExpressionDialect)
	runLex(l)
	return l.NextToken()
}

func verifyExprTokens(t *testing.T, expString string, tokens []Token) {
	l := NewLexer(expString, ExpressionDialect)
	for _, goodToken := range tokens {
		tok := l.NextToken()
		//u.Debugf("%#v  %#v", tok, goodToken)
		assert.Equalf(t, tok.T, goodToken.T, "want='%v' has %v ", goodToken.T, tok.T)
		assert.Equalf(t, tok.V, goodToken.V, "want='%v' has %v ", goodToken.V, tok.V)
	}
}
func TestLexExprDialect(t *testing.T) {
	verifyExprTokens(t, `eq(toint(item),5)`,
		[]Token{
			tv(TokenUdfExpr, "eq"),
			tv(TokenLeftParenthesis, "("),
			tv(TokenUdfExpr, "toint"),
			tv(TokenLeftParenthesis, "("),
			tv(TokenIdentity, "item"),
			tv(TokenRightParenthesis, ")"),
			tv(TokenComma, ","),
			tv(TokenInteger, "5"),
			tv(TokenRightParenthesis, ")"),
		})
}
