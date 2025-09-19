package repl

import (
	"bufio"
	"fmt"
	"io"
	"pika/lexer"
	"pika/parser"
)

const PROMPT = ">>"

// func to start the repl
func Start(in io.Reader, out io.Writer) {
	// create scanner to take in user input
	scanner := bufio.NewScanner(in)

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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors[]string) {
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}
