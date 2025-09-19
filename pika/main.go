package main

import (
	"fmt"
	"os"
	"os/user"
	"pika/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(repl.PIKA)
	fmt.Printf("Hello %s! This is the PIKA programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands!\n\n")
	repl.Start(os.Stdin, os.Stdout)
}
