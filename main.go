package main

import (
	"fmt"
	"os"

	"github.com/Krawabbel/go-lox/lox"
)

func main() {

	switch len(os.Args) {

	case 1:
		if err := lox.RunREPL(); err != nil {
			panic(err)
		}

	case 2:

		arg := os.Args[1]

		if err := lox.RunScript(arg); err != nil {
			panic(err)
		}

	default:
		fmt.Println("usage: lox [path/to/script.lox]")

	}

}
