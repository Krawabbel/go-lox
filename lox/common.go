package lox

import (
	"fmt"
	"io"
	"os"
)

var OUTPUT io.Writer = os.Stdout
var INPUT io.Reader = os.Stdin

func print(format string, a ...any) {
	fmt.Fprintf(OUTPUT, format, a...)
}
