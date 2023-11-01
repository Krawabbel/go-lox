package lox

type Scanner struct {
	start, curr int
	code        string
	line        int
}

func MakeScanner(src string) Scanner {
	return Scanner{code: src, line: 1}
}

func (scanner Scanner) check_eof(pos int) bool {
	return pos >= len(scanner.code)
}

func (scanner Scanner) eof() bool {
	return scanner.check_eof(scanner.curr)
}

func (scanner Scanner) peek_curr() byte {
	return scanner.code[scanner.curr]
}

func (scanner Scanner) peek_next() byte {
	return scanner.code[scanner.curr+1]
}

func (scanner *Scanner) scan_byte() byte {
	var b = scanner.code[scanner.curr]
	scanner.curr++
	return b
}

func (scanner *Scanner) advance() {
	scanner.curr++
}

func (scanner *Scanner) scan_string() Token {

	for !scanner.eof() {

		switch scanner.peek_curr() {

		case '\n':
			scanner.line++

		case '"':
			scanner.advance() // the closing quote
			return scanner.make_token(TOKEN_STRING)
		}

		scanner.advance()
	}

	return scanner.error_token("unterminated string")

}

func (scanner *Scanner) scan_number() Token {

	for scanner.check_digit() {
		scanner.advance()
	}

	if scanner.match_byte('.') {
		if !scanner.check_digit() {
			return scanner.error_token("unterminated floating point number")
		}
		scanner.advance()
	}

	for scanner.check_digit() {
		scanner.advance()
	}

	return scanner.make_token(TOKEN_NUMBER)
}

func (scanner *Scanner) scan_identifier() Token {

	for scanner.check_alpha() || scanner.check_digit() {
		scanner.advance()
	}
	return scanner.make_token(scanner.identifier_spec())
}

func (scanner Scanner) check_alpha() bool {
	if scanner.eof() {
		return false
	}
	return is_alpha(scanner.peek_curr())
}

func (scanner *Scanner) check_digit() bool {
	if scanner.eof() {
		return false
	}
	return is_digit(scanner.peek_curr())
}

func (scanner *Scanner) match_byte(b byte) bool {
	if scanner.eof() {
		return false
	}

	if scanner.peek_curr() != b {
		return false
	}

	scanner.advance()
	return true
}

func (scanner *Scanner) skip_whitespace() {
	for !scanner.eof() {
		switch scanner.peek_curr() {
		case ' ', '\r', '\t':
			scanner.advance()
		case '\n':
			scanner.line++
			scanner.advance()
		case '/':
			if scanner.peek_next() == '/' {
				for !scanner.eof() && scanner.peek_curr() != '\n' {
					scanner.advance()
				}
			} else {
				return // we are actually in the middle of a division
			}
		default:
			return
		}
	}
}

func (scanner *Scanner) scan_token() Token {

	scanner.skip_whitespace()

	scanner.start = scanner.curr

	if scanner.eof() {
		return scanner.eof_token()
	}

	var c = scanner.scan_byte()

	if is_alpha(c) {
		return scanner.scan_identifier()
	}

	if is_digit(c) {
		return scanner.scan_number()
	}

	switch c {
	case '(':
		return scanner.make_token(TOKEN_LEFT_PAREN)
	case ')':
		return scanner.make_token(TOKEN_RIGHT_PAREN)
	case '{':
		return scanner.make_token(TOKEN_LEFT_BRACE)
	case '}':
		return scanner.make_token(TOKEN_RIGHT_BRACE)
	case ';':
		return scanner.make_token(TOKEN_SEMICOLON)
	case ',':
		return scanner.make_token(TOKEN_COMMA)
	case '.':
		return scanner.make_token(TOKEN_DOT)
	case '-':
		return scanner.make_token(TOKEN_MINUS)
	case '+':
		return scanner.make_token(TOKEN_PLUS)
	case '/':
		return scanner.make_token(TOKEN_SLASH)
	case '*':
		return scanner.make_token(TOKEN_STAR)

	case '!':
		return scanner.make_token(if_then_else(scanner.match_byte('='), TOKEN_BANG_EQUAL, TOKEN_BANG))
	case '=':
		return scanner.make_token(if_then_else(scanner.match_byte('='), TOKEN_EQUAL_EQUAL, TOKEN_EQUAL))
	case '<':
		return scanner.make_token(if_then_else(scanner.match_byte('='), TOKEN_LESS_EQUAL, TOKEN_LESS))
	case '>':
		return scanner.make_token(if_then_else(scanner.match_byte('='), TOKEN_GREATER_EQUAL, TOKEN_GREATER))

	case '"':
		return scanner.scan_string()
	}

	return scanner.error_token("unexpected character")
}

func (scanner Scanner) make_token(spec int) Token {
	return Token{
		spec:   spec,
		lexeme: scanner.code[scanner.start:scanner.curr],
		line:   scanner.line,
	}
}

func (scanner Scanner) error_token(msg string) Token {
	return Token{
		spec:   TOKEN_ERROR,
		lexeme: msg,
		line:   scanner.line,
	}
}

func (scanner Scanner) eof_token() Token {
	return Token{
		spec:   TOKEN_EOF,
		lexeme: "EOF",
		line:   scanner.line,
	}
}

func (scanner Scanner) identifier_spec() int {

	switch scanner.code[scanner.start] {

	case 'a':
		return scanner.check_keyword(1, 2, "nd", TOKEN_AND)

	case 'c':
		return scanner.check_keyword(1, 4, "lass", TOKEN_CLASS)

	case 'e':
		return scanner.check_keyword(1, 3, "lse", TOKEN_ELSE)

	case 'f':

		if len(scanner.code) > 1 {

			switch scanner.code[scanner.start+1] {

			case 'a':
				return scanner.check_keyword(2, 3, "lse", TOKEN_FALSE)

			case 'o':
				return scanner.check_keyword(2, 1, "r", TOKEN_FOR)

			case 'u':
				return scanner.check_keyword(2, 1, "n", TOKEN_FUN)
			}
		}

	case 'i':
		return scanner.check_keyword(1, 1, "f", TOKEN_IF)

	case 'n':
		return scanner.check_keyword(1, 2, "il", TOKEN_NIL)

	case 'o':
		return scanner.check_keyword(1, 1, "r", TOKEN_OR)

	case 'p':
		return scanner.check_keyword(1, 4, "rint", TOKEN_PRINT)

	case 'r':
		return scanner.check_keyword(1, 5, "eturn", TOKEN_RETURN)

	case 's':
		return scanner.check_keyword(1, 4, "uper", TOKEN_SUPER)

	case 't':

		if len(scanner.code) > 1 {

			switch scanner.code[scanner.start+1] {

			case 'h':
				return scanner.check_keyword(2, 2, "is", TOKEN_THIS)

			case 'r':
				return scanner.check_keyword(2, 2, "ue", TOKEN_TRUE)
			}
		}

	case 'v':
		return scanner.check_keyword(1, 2, "ar", TOKEN_VAR)

	case 'w':
		return scanner.check_keyword(1, 4, "hile", TOKEN_WHILE)

	}

	return TOKEN_IDENTIFIER
}

func (scanner Scanner) check_keyword(start, length int, rem string, spec int) int {

	var match = (scanner.curr-scanner.start == start+length) && (scanner.code[scanner.start+start:scanner.start+start+length] == rem)

	return if_then_else(match, spec, TOKEN_IDENTIFIER)
}

func is_digit(b byte) bool {
	return '0' <= b && b <= '9'
}

func is_alpha(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || (b == '_')
}
