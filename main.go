package main

import (
	"Ape/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Ape programming language!\n", user)
	fmt.Printf("Feel free to type in commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
