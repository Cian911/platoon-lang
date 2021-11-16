package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/cian911/platoon/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Thank you for using the Platton programming language.\n", user.Username)

	// Start the REPL
	repl.Start(os.Stdin, os.Stdout)
}
