package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("HELLO " + user.Username)
	fmt.Println(os.Getenv("USER"))
	fmt.Println("Let's have a fun time with the Monkey language!")
	repl.Start(os.Stdin, os.Stdout)
}
