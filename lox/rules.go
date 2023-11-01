package lox

func get_rule(p *Parser, spec int) (prefix, infix func(), prec int) {

	switch spec {

	case TOKEN_LEFT_PAREN:
		return p.parse_grouping, nil, PREC_NONE

	case TOKEN_RIGHT_PAREN:
		return nil, nil, PREC_NONE

	case TOKEN_LEFT_BRACE:
		return nil, nil, PREC_NONE

	case TOKEN_RIGHT_BRACE:
		return nil, nil, PREC_NONE

	case TOKEN_COMMA:
		return nil, nil, PREC_NONE

	case TOKEN_DOT:
		return nil, nil, PREC_NONE

	case TOKEN_MINUS:
		return p.parse_unary, p.parse_binary, PREC_TERM

	case TOKEN_PLUS:
		return nil, p.parse_binary, PREC_TERM

	case TOKEN_SEMICOLON:
		return nil, nil, PREC_NONE

	case TOKEN_SLASH:
		return nil, p.parse_binary, PREC_FACTOR

	case TOKEN_STAR:
		return nil, p.parse_binary, PREC_FACTOR

	case TOKEN_BANG:
		return p.parse_unary, nil, PREC_NONE

	case TOKEN_BANG_EQUAL:
		return nil, p.parse_binary, PREC_EQUALITY

	case TOKEN_EQUAL:
		return nil, nil, PREC_NONE

	case TOKEN_EQUAL_EQUAL:
		return nil, p.parse_binary, PREC_EQUALITY

	case TOKEN_GREATER:
		return nil, p.parse_binary, PREC_COMPARISON

	case TOKEN_GREATER_EQUAL:
		return nil, p.parse_binary, PREC_COMPARISON

	case TOKEN_LESS:
		return nil, p.parse_binary, PREC_COMPARISON

	case TOKEN_LESS_EQUAL:
		return nil, p.parse_binary, PREC_COMPARISON

	case TOKEN_IDENTIFIER:
		return nil, nil, PREC_NONE

	case TOKEN_STRING:
		return p.parse_string, nil, PREC_NONE

	case TOKEN_NUMBER:
		return p.parse_number, nil, PREC_NONE

	case TOKEN_AND:
		return nil, nil, PREC_NONE

	case TOKEN_CLASS:
		return nil, nil, PREC_NONE

	case TOKEN_ELSE:
		return nil, nil, PREC_NONE

	case TOKEN_FALSE:
		return p.parse_literal, nil, PREC_NONE

	case TOKEN_FOR:
		return nil, nil, PREC_NONE

	case TOKEN_FUN:
		return nil, nil, PREC_NONE

	case TOKEN_IF:
		return nil, nil, PREC_NONE

	case TOKEN_NIL:
		return p.parse_literal, nil, PREC_NONE

	case TOKEN_OR:
		return nil, nil, PREC_NONE

	case TOKEN_PRINT:
		return nil, nil, PREC_NONE

	case TOKEN_RETURN:
		return nil, nil, PREC_NONE

	case TOKEN_SUPER:
		return nil, nil, PREC_NONE

	case TOKEN_THIS:
		return nil, nil, PREC_NONE

	case TOKEN_TRUE:
		return p.parse_literal, nil, PREC_NONE

	case TOKEN_VAR:
		return nil, nil, PREC_NONE

	case TOKEN_WHILE:
		return nil, nil, PREC_NONE

	case TOKEN_ERROR:
		return nil, nil, PREC_NONE

	case TOKEN_EOF:
		return nil, nil, PREC_NONE
	}

	panic("unexpected token")
}
