package lexer

import (
	"pika/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // char after current position
	ch           byte // char under examination
}

// create new lexer
func New(input string) *Lexer {
	l := &Lexer{input: input} // initialize it with input string
	l.readChar()              // set current char and advance the lexer position
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // check if at end of input and set char to 0 (ascii for nul)
	} else {
		l.ch = l.input[l.readPosition] // else set char from next position
	}

	l.position = l.readPosition // increment both positions
	l.readPosition += 1
}

// returns token depending on which char it
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	// don't count whitespaces
	l.skipWhitespace();

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if l.isLetter(l.ch) {
			// check if char is letter, if yes, read whole literal
			// if literal is not keyword return IDENT else return keyword
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch){
			// check if digit, return type and literal accordingly
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// creates a Token from tokenType and ch
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string{
	position := l.position
	for l.isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position: l.position]
}

func (l *Lexer) isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

// skip any white spaces in the input
func (l *Lexer) skipWhitespace() {
	if l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// read a number like 123 by looping through each char in string until 
// a non digit char is found and return the number
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position: l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}