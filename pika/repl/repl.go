package repl

import (
	"bufio"
	"fmt"
	"io"
	"pika/evaluator"
	"pika/lexer"
	"pika/object"
	"pika/parser"
)

const PROMPT = ">> "

const PIKA = `
       .__ __            
______ |__|  | _______   
\____ \|  |  |/ /\__  \  
|  |_> >  |    <  / __ \_
|   __/|__|__|_ \(____  /
|__|           \/     \/ 

`

// func to start the repl
func Start(in io.Reader, out io.Writer) {
	// create scanner to take in user input
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	// infinite read eval print loop
	for {
		// print >>
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		// if nothing scanned return
		if !scanned {
			return
		}
		// get the line and initialize a lexer with the line as input
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "The parser flinched! invalid move detected.\n")
	io.WriteString(out, " Error log:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
