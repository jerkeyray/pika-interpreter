package evaluator

import (
	"pika/ast"
	"pika/object"
)

// eval takes in an ast node and returns appropriate object 
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	// eval statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// eval expressions
	case *ast.IntegerLiteral: 
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

