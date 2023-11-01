package lox

import "fmt"

func compile(src string) (*Chunk, bool) {

	var parser Parser

	parser.scanner = MakeScanner(src)

	// parser.had_error = false
	// parser.is_in_panic_mode = false

	parser.step()

	parser.parse_expression()

	parser.consume(TOKEN_EOF, "expected end of expression")

	parser.end_compiler()

	if DEBUG_PRINT_CODE && !parser.had_error {
		fmt.Fprintln(STDDBG, disassemble_chunk(parser.current_chunk(), "code"))
	}

	return parser.current_chunk(), !parser.had_error

}
