package lox

import "testing"

func TestScanner(t *testing.T) {

	var args = []struct {
		given      string
		want       int
		want_given bool
	}{
		{"(", TOKEN_LEFT_PAREN, true},
		{")", TOKEN_RIGHT_PAREN, true},
		{"{", TOKEN_LEFT_BRACE, true},
		{"}", TOKEN_RIGHT_BRACE, true},
		{",", TOKEN_COMMA, true},
		{".", TOKEN_DOT, true},
		{"-", TOKEN_MINUS, true},
		{"+", TOKEN_PLUS, true},
		{";", TOKEN_SEMICOLON, true},
		{"/", TOKEN_SLASH, true},
		{"*", TOKEN_STAR, true},
		{"!", TOKEN_BANG, true},
		{"!=", TOKEN_BANG_EQUAL, true},
		{"=", TOKEN_EQUAL, true},
		{"==", TOKEN_EQUAL_EQUAL, true},
		{">", TOKEN_GREATER, true},
		{">=", TOKEN_GREATER_EQUAL, true},
		{"<", TOKEN_LESS, true},
		{"<=", TOKEN_LESS_EQUAL, true},
		{"and", TOKEN_AND, true},
		{"class", TOKEN_CLASS, true},
		{"else", TOKEN_ELSE, true},
		{"false", TOKEN_FALSE, true},
		{"for", TOKEN_FOR, true},
		{"fun", TOKEN_FUN, true},
		{"if", TOKEN_IF, true},
		{"nil", TOKEN_NIL, true},
		{"or", TOKEN_OR, true},
		{"print", TOKEN_PRINT, true},
		{"return", TOKEN_RETURN, true},
		{"super", TOKEN_SUPER, true},
		{"this", TOKEN_THIS, true},
		{"true", TOKEN_TRUE, true},
		{"var", TOKEN_VAR, true},
		{"while", TOKEN_WHILE, true},

		{"0", TOKEN_NUMBER, true},
		{"0.0", TOKEN_NUMBER, true},
		{"00", TOKEN_NUMBER, true},
		{"0.00", TOKEN_NUMBER, true},

		{"0.", TOKEN_ERROR, false},

		{"a", TOKEN_IDENTIFIER, true},
		{"aa", TOKEN_IDENTIFIER, true},
		{"\"\"", TOKEN_STRING, true},
		{"\"b\"", TOKEN_STRING, true},
		{"\"bb\"", TOKEN_STRING, true},

		{"~", TOKEN_ERROR, false},

		{"", TOKEN_EOF, false},
	}

	var src = "\t"
	for _, arg := range args {
		src += arg.given + " \n \t \n"
	}

	var scanner = MakeScanner(src)

	for _, arg := range args {

		t.Log(arg.given)

		if scanner.eof() {
			t.Fatalf("unexpected EOF")
		}

		var have = scanner.scan_token()

		if (arg.want_given) && (have.lexeme != arg.given) {
			t.Fatalf("scan_token(): lexeme: want %v, have %v", arg.given, have.lexeme)
		}

		if have.spec != arg.want {
			t.Fatalf("scan_token(): spec: want %v, have %v", arg.want, have.spec)
		}

	}

	if !scanner.eof() {
		t.Fatalf("expected EOF")
	}

}
