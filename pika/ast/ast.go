package ast

import "go/token"

// top level design
// ast program -> statements -> let statement -> identifier(name) and expression(value)

// interface for every node in our ast
type Node interface {
	TokenLiteral() string
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
func(p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// for say x = 5, x is identifier and 5 is value(expression)
type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value	Expression
}

// methods to satisfy statement interface
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() {return ls.Token.Literal}

// identifier should only have the value
type Identifier struct {
	Token token.Token
	Value string
}

// methods to satisfy expression interface since it can produce values in some cases
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() {return i.Token.Literal}