package lox

import (
	"io"
	"os"
)

var STDOUT io.Writer = os.Stdout
var STDERR io.Writer = os.Stderr

var STDDBG io.Writer = os.Stdout

var STDIN io.Reader = os.Stdin
