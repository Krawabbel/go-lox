package lox

const (
	// Single-character tokens.
	TOKEN_LEFT_PAREN = iota
	TOKEN_RIGHT_PAREN
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_COMMA
	TOKEN_DOT
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_SEMICOLON
	TOKEN_SLASH
	TOKEN_STAR
	// One or two character tokens.
	TOKEN_BANG
	TOKEN_BANG_EQUAL
	TOKEN_EQUAL
	TOKEN_EQUAL_EQUAL
	TOKEN_GREATER
	TOKEN_GREATER_EQUAL
	TOKEN_LESS
	TOKEN_LESS_EQUAL
	// Literals.
	TOKEN_IDENTIFIER
	TOKEN_STRING
	TOKEN_NUMBER
	// Keywords.
	TOKEN_AND
	TOKEN_CLASS
	TOKEN_ELSE
	TOKEN_FALSE
	TOKEN_FOR
	TOKEN_FUN
	TOKEN_IF
	TOKEN_NIL
	TOKEN_OR
	TOKEN_PRINT
	TOKEN_RETURN
	TOKEN_SUPER
	TOKEN_THIS
	TOKEN_TRUE
	TOKEN_VAR
	TOKEN_WHILE
	// Special.
	TOKEN_ERROR
	TOKEN_EOF
)

type Token struct {
	typ    int
	lexeme string
	line   int
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

type Scanner struct {
	start, curr, line int
	code              string
}

func (scanner *Scanner) scan() Token {
	scanner.start = scanner.curr
	if scanner.eof() {
		return scanner.make_token(TOKEN_EOF)
	}

	var c = scanner.advance()

	switch c {

	}

	return scanner.error_token("unexpected character")
}

func (scanner Scanner) eof() bool {
	return scanner.curr == len(scanner.code)
}

func NewScanner(src string) *Scanner {
	return &Scanner{code: src, line: 1}
}
