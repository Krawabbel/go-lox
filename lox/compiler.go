package lox

import "fmt"

func compile(src string) error {
	var scanner = NewScanner(src)
	var line = -1
	for {
		var token = scanner.scan()
		if token.line != line {
			print(fmt.Sprintf("%4d", token.line))
			line = token.line
		} else {
			print("   | ")
		}
		print(fmt.Sprintf("%2d '%.*s'\n", token.typ, len(token.lexeme), token.lexeme))

		if token.typ == TOKEN_EOF {
			break
		}
	}

	return nil
}
