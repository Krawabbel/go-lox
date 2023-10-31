package lox

func is_digit(b byte) bool {
	return '0' <= b && b <= '9'
}

func is_alpha(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || (b == '_')
}

type Scanner struct {
	start, curr, line int
	code              string
}

func NewScanner(src string) *Scanner {
	return &Scanner{code: src, line: 1}
}

func (scanner Scanner) eof() bool {
	return scanner.curr == len(scanner.code)
}

func (scanner Scanner) peek_curr() byte {
	return scanner.code[scanner.curr]
}

func (scanner Scanner) peek_next() byte {
	if scanner.eof() {
		return 0
	}
	return scanner.code[scanner.curr+1]
}

func (scanner *Scanner) advance() {
	scanner.curr++
}

func (scanner *Scanner) scan_byte() byte {
	var b = scanner.peek_curr()
	scanner.advance()
	return b
}

func (scanner *Scanner) scan_string() Token {

	for !scanner.eof() && scanner.peek_curr() != '"' {
		if scanner.peek_curr() == '\n' {
			scanner.line++
		}
		scanner.advance()
	}

	if scanner.eof() {
		return scanner.error_token("unterminated string")
	}

	scanner.advance() // the closing quote

	return scanner.make_token(TOKEN_STRING)
}

func (scanner *Scanner) scan_number() Token {

	for is_digit(scanner.peek_curr()) {
		scanner.advance()
	}

	if (scanner.peek_curr() == '.') && is_digit(scanner.peek_next()) {
		scanner.advance()
	}

	for is_digit(scanner.peek_curr()) {
		scanner.advance()
	}

	return scanner.make_token(TOKEN_NUMBER)
}

func (scanner *Scanner) scan_identifier() Token {
	for is_alpha(scanner.peek_curr()) || is_digit(scanner.peek_curr()) {
		scanner.advance()
	}
	return scanner.make_token(TOKEN_IDENTIFIER)
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
				for scanner.peek_curr() != '\n' && !scanner.eof() {
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

func (scanner Scanner) make_token(typ int) Token {
	return Token{
		typ:    typ,
		lexeme: scanner.code[scanner.start:scanner.curr],
		line:   scanner.line,
	}
}

func (scanner Scanner) error_token(msg string) Token {
	return Token{
		typ:    TOKEN_ERROR,
		lexeme: msg,
		line:   scanner.line,
	}
}

func (scanner Scanner) eof_token() Token {
	return Token{
		typ:    TOKEN_EOF,
		lexeme: "EOF",
		line:   scanner.line,
	}
}
