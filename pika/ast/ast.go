package ast

import (
	"pika/token"
	"bytes"
)

// top level design
// ast program -> statements -> let statement -> identifier(name) and expression(value)

// interface for every node in our ast
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// method on program to return token literal of first statement
// if not found return empty string
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// for say x = 5, x is identifier and 5 is value(expression)
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// methods to satisfy statement interface
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// identifier should only have the value
type Identifier struct {
	Token token.Token
	Value string
}

// methods to satisfy expression interface since it can produce values in some cases
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// return statement 
type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {return rs.Token.Literal}

// implement expression statements
type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {return es.Token.Literal}

// create a buffer and write value of each statments String() method to it
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}


