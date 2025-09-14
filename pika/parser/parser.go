package parser

import (
	"pika/ast"
	"pika/token"
	"pika/lexer"
)

type Parser struct {
	// pointer to our lexer which contniuously reads chars
	l *lexer.Lexer
	// our current token and next token just like chars in the lexer
	curToken	token.Token
	peekToken	token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens to set curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// helper to read next token from the lexer
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}