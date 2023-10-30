package lox

import (
	"fmt"
	"io"
	"os"
)

var OUTPUT io.Writer = os.Stdout
var INPUT io.Reader = os.Stdin

func print_ln(msg string) {
	fmt.Fprintln(OUTPUT, msg)
}

func print(msg string) {
	fmt.Fprint(OUTPUT, msg)
}
