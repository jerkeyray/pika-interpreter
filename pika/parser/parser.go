package parser

import (
	"pika/ast"
	"pika/lexer"
	"pika/token"
	"fmt"
)

type Parser struct {
	// pointer to our lexer which contniuously reads chars
	l *lexer.Lexer
	// our current token and next token just like chars in the lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

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
	// root node of ast
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// iterate until eof token
	// parse each statement and if not nil add to statments slice
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// parse each statement based on token type
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// method to be called by parseStatement() 
// first create *ast.LetStatement node
// use ident token to create identifier node
// look for assign token then jumps over the expression from equal sign to semicolon
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// compare current token to given token
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// compare next to token to given token
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// check the type of peekToken and only if the type is correct do we move to next token
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}
 
// add error msg to error field in parser when peek token doesn't match token type required
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}