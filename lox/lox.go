package lox

import (
	"bufio"
	"fmt"
	"os"
)

func RunREPL() error {

	var input = bufio.NewScanner(INPUT)

	var vm = new(VM)
	for print("> "); input.Scan(); print("> ") {

		var line = input.Text()
		if line == "exit" {
			break
		}
		vm.interpret(line)
	}
	return nil
}

func RunFile(path string) error {

	src, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var vm = new(VM)

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
