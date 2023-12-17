package main

import (
	"fmt"
	"go-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome, %s! This an c like interpreter repl \n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
