package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Username)
	fmt.Println(os.Getenv("USER"))
}
