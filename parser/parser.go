package parser

import (
	"Ape/lexer"
	"Ape/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{
		l: l,
	}
}
