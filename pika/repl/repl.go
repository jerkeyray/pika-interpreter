package repl

import (
	"bufio"
	"io"
	"fmt"
	"pika/lexer"
	"pika/token"
)

const PROMPT = ">>"

// func to start the repl
func Start(in io.Reader, out io.Reader) {
	// create scanner to take in user input
	scanner := bufio.NewScanner(in)

	// infinite read eval print loop
	for {
		// print >>
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		// if nothing scanned return 
		if !scanned {
			return 
		}
		// get the line and initialize a lexer with the line as input
		line := scanner.Text()
		l := lexer.New(line)
		
		// print all the tokens from the lexer
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}