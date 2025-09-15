package parser

import (
	"fmt"
	"pika/ast"
	"pika/lexer"
	"pika/token"
	"strconv"
)

type Parser struct {
	// pointer to our lexer which contniuously reads chars
	l *lexer.Lexer
	// our current token and next token just like chars in the lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string

	// use map to check if there is a parsing function associated with current token type
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// initialize parseFns maps on parser and register a parsing function
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	// read two tokens to set curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// return ast.Identifer with current token and its value
// does not advance the tokens
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
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
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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

// build ast node for return statement with the current token
// then bring parser in place for expression by calling nextToken until it finds a semicolon
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression // this takes left expression as an arguement
)

// helper methods to add entries to the fn maps
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	// build out ast node for expression statements
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	// semicolon is optional in this statement
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// precedence order of pika programming language
const (
	_ int = iota // gives following constants incrementing numbers as values (1 - 7)
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	// check if any parse fn is associated with current token type
	// if yes call it
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()
	return leftExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	// convert string to int64
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as an integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	// return newly constructed ast.IntegerLiteral node with int64 value
	lit.Value = value

	return lit
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg :=	fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	// build an ast node and advance the token
	expression := &ast.PrefixExpression{
		Token: p.curToken,
		Operator: p.curToken.Literal, 
	}

	p.nextToken()
	// call parse expression again after advancing the token
	expression.Right = p.parseExpression(PREFIX)

	return expression
}