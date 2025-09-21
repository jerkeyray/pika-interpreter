package evaluator

import (
	"fmt"
	"pika/object"
)

// map for string names to actual functions
var builtins = map[string]*object.Builtin{
	// map "len" to new builtin function object
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// check number of arguments
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			// check argument type
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	}, 

	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// check number of arguments
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			// check argument type
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}
			// return first element of array else null if empty
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return NULL
		},
	},

	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// check number of arguments
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			// check argument type
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}
			// return last element of array else null if empty
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[len(arr.Elements)-1]
			}
			return NULL
		},
	},

	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// check number of arguments
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			// check argument type
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}
			// return array containing all elements except first else null if empty
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length)
				copy(newElements, arr.Elements[1: length]) // copy all elements except first
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},

	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// check number of arguments
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			// check first argument type
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `push` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements) // copy all elements of original array
			newElements[length] = args[1]   // add new element at the end

			return &object.Array{Elements: newElements}
		},
	},

	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				// print the string representation of each argument
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}