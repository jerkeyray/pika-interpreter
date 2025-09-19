package object

import (
	"fmt"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
)

// whenever we encounter an integer in source code
//  we change to ast.IntegerLiteral and while evaluating AST node
// convert it to Object.Integer saving the value in a struct and passing its reference
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {return INTEGER_OBJ}
func (i *Integer) Inspect() string {return fmt.Sprintf("%d", i.Value)}

// boolean
type Boolean struct {
	Value bool 
}

func (b *Boolean) Type() ObjectType {return BOOLEAN_OBJ}
func (b *Boolean) Inspect() string {return fmt.Sprintf("%t", b.Value)}

// null
type Null struct {}

func (n *Null) Type() ObjectType {return NULL_OBJ}
func (n *Null) Inspect() string {return "null"}