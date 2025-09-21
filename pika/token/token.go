package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	BANG     = "!"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	// Grouping
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET	= "["
	RBRACKET	= "]"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	STRING   = "STRING"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	// look for keyowrd in map and return it
	// if not found return IDENT
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
