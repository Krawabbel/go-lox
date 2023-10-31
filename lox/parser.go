package lox

import (
	"fmt"
	"strconv"
)

const (
	PREC_NONE       = iota
	PREC_ASSIGNMENT // =
	PREC_OR         // or
	PREC_AND        // and
	PREC_EQUALITY   // == !=
	PREC_COMPARISON // < > <= >=
	PREC_TERM       // + -
	PREC_FACTOR     // * /
	PREC_UNARY      // ! -
	PREC_CALL       // . ()
	PREC_PRIMARY
)

type Parser struct {
	curr_token, prev_token Token
	scanner                Scanner

	compiling_chunk Chunk

	had_error, is_in_panic_mode bool
}

func (parser *Parser) step() {
	parser.prev_token = parser.curr_token

	for {
		parser.curr_token = parser.scanner.scan_token()
		if parser.curr_token.spec != TOKEN_ERROR {
			break
		}
		parser.report_error_at_current(parser.curr_token.lexeme)
	}
}

func (parser *Parser) consume(spec int, msg string) {
	if parser.curr_token.spec == spec {
		parser.step()
		return
	}
	parser.report_error_at_current(msg)
}

func (parser Parser) current_chunk() *Chunk {
	return &parser.compiling_chunk
}

func (parser *Parser) end_compiler() {
	parser.emit_return()
}

func (parser *Parser) parse_expression() {
	parser.parse_precedence(PREC_ASSIGNMENT)
}

func (parser *Parser) parse_number() {
	var value, err = strconv.ParseFloat(parser.prev_token.lexeme, 64)
	if err != nil {
		parser.report_error(err.Error())
	}
	parser.emit_constant(value)
}

func (parser *Parser) parse_grouping() {
	parser.parse_expression()
	parser.consume(TOKEN_RIGHT_PAREN, "expect ')' after expression")
}

func (parser *Parser) parse_unary() {

	var operator_spec = parser.prev_token.spec

	parser.parse_precedence(PREC_UNARY)

	switch operator_spec {
	case TOKEN_MINUS:
		parser.emit_byte(OP_NEGATE)

	default:
		return // unreachable
	}
}

func (parser *Parser) parse_binary() {

	var operator_spec = parser.prev_token.spec

	var _, _, prec = parser.get_rule(operator_spec)
	parser.parse_precedence(prec + 1)

	switch operator_spec {
	case TOKEN_PLUS:
		parser.emit_byte(OP_ADD)
	case TOKEN_MINUS:
		parser.emit_byte(OP_SUBTRACT)
	case TOKEN_STAR:
		parser.emit_byte(OP_MULTIPLY)
	case TOKEN_SLASH:
		parser.emit_byte(OP_DIVIDE)
	default:
		//unreachable
	}

}

func (parser *Parser) parse_precedence(prec int) {

	parser.step()
	var prefix_rule, _, _ = parser.get_rule(parser.prev_token.spec)

	if prefix_rule == nil {
		parser.report_error("expect expression")
		return
	}

	prefix_rule()

	for _, _, curr_prec := parser.get_rule(parser.curr_token.spec); prec <= curr_prec; _, _, curr_prec = parser.get_rule(parser.curr_token.spec) {
		parser.step()
		var _, infix_rule, _ = parser.get_rule(parser.prev_token.spec)
		infix_rule()
	}
}

func (parser *Parser) get_rule(spec int) (prefix, infix func(), prec int) {

	switch prec {

	case TOKEN_LEFT_PAREN:
		return parser.parse_grouping, nil, PREC_NONE

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
		return parser.parse_unary, parser.parse_binary, PREC_TERM

	case TOKEN_PLUS:
		return nil, parser.parse_binary, PREC_TERM

	case TOKEN_SEMICOLON:
		return nil, nil, PREC_NONE

	case TOKEN_SLASH:
		return nil, parser.parse_binary, PREC_FACTOR

	case TOKEN_STAR:
		return nil, parser.parse_binary, PREC_FACTOR

	case TOKEN_BANG:
		return nil, nil, PREC_NONE

	case TOKEN_BANG_EQUAL:
		return nil, nil, PREC_NONE

	case TOKEN_EQUAL:
		return nil, nil, PREC_NONE

	case TOKEN_EQUAL_EQUAL:
		return nil, nil, PREC_NONE

	case TOKEN_GREATER:
		return nil, nil, PREC_NONE

	case TOKEN_GREATER_EQUAL:
		return nil, nil, PREC_NONE

	case TOKEN_LESS:
		return nil, nil, PREC_NONE

	case TOKEN_LESS_EQUAL:
		return nil, nil, PREC_NONE

	case TOKEN_IDENTIFIER:
		return nil, nil, PREC_NONE

	case TOKEN_STRING:
		return nil, nil, PREC_NONE

	case TOKEN_NUMBER:
		return parser.parse_number, nil, PREC_NONE

	case TOKEN_AND:
		return nil, nil, PREC_NONE

	case TOKEN_CLASS:
		return nil, nil, PREC_NONE

	case TOKEN_ELSE:
		return nil, nil, PREC_NONE

	case TOKEN_FALSE:
		return nil, nil, PREC_NONE

	case TOKEN_FOR:
		return nil, nil, PREC_NONE

	case TOKEN_FUN:
		return nil, nil, PREC_NONE

	case TOKEN_IF:
		return nil, nil, PREC_NONE

	case TOKEN_NIL:
		return nil, nil, PREC_NONE

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
		return nil, nil, PREC_NONE

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

func (parser *Parser) emit_byte(bs ...byte) {
	for _, b := range bs {
		parser.current_chunk().write(b, parser.prev_token.line)
	}
}

func (parser *Parser) emit_return() {
	parser.emit_byte(OP_RETURN)
}

func (parser *Parser) emit_constant(value Value) {
	parser.emit_byte(OP_CONSTANT, parser.make_constant(value))
}

func (parser *Parser) make_constant(value Value) byte {
	var addr = parser.current_chunk().add_constant(value)
	if addr > 0xFF {
		panic("too many constants in one chunk")
	}
	return byte(addr)
}

func (parser *Parser) report_error_at_current(msg string) {
	parser.report_error_at(parser.curr_token, msg)
}

func (parser *Parser) report_error(msg string) {
	parser.report_error_at(parser.prev_token, msg)
}

func (parser *Parser) report_error_at(token Token, msg string) {

	if parser.is_in_panic_mode {
		return
	}
	parser.is_in_panic_mode = true

	fmt.Fprintf(STDERR, "[line %d] error", token.line)

	switch token.spec {

	case TOKEN_EOF:
		fmt.Fprintf(STDERR, " at end")
	case TOKEN_ERROR:
		// nothing
	default:
		fmt.Fprintf(STDERR, " at '%s'", token.lexeme)
	}

	fmt.Fprintf(STDERR, ": %s\n", msg)

	parser.had_error = true
}
