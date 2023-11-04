package lox

import (
	"bufio"
	"fmt"
	"os"
)

func RunREPL() error {

	var input = bufio.NewScanner(STDIN)

	var vm = newVM()

	for fmt.Fprintf(STDOUT, "> "); input.Scan(); fmt.Fprintf(STDOUT, "> ") {

		var line = input.Text()
		switch line {
		case "exit":
			return nil
		case "":
			// do nothing
		default:
			vm.interpret(line)
		}
	}

	return nil
}

func RunScript(path string) error {

	src, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var vm = newVM()

	var interpreter_result = vm.interpret(string(src))
	switch interpreter_result {
	case INTERPRET_OK:
		// do nothing
	case INTERPRET_RUNTIME_ERROR:
		return fmt.Errorf("runtime error")
	case INTERPRET_COMPILE_ERROR:
		return fmt.Errorf("compile error")
	default:
		return fmt.Errorf("unexpected interpreter result code: %v", interpreter_result)
	}

	return nil
}
