package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/KasonBraley/monkey-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("type in commands")

	repl.Start(os.Stdin, os.Stdout)
}
