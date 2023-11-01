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

func (parser *Parser) current_chunk() *Chunk {
	return &parser.compiling_chunk
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

func (parser *Parser) end_compiler() {
	parser.emit_return()
}

func (parser *Parser) parse_expression() {
	parser.parse_precedence(PREC_ASSIGNMENT)
}

func (parser *Parser) parse_literal() {
	switch parser.prev_token.spec {
	case TOKEN_FALSE:
		parser.emit_byte(OP_FALSE)
	case TOKEN_NIL:
		parser.emit_byte(OP_NIL)
	case TOKEN_TRUE:
		parser.emit_byte(OP_TRUE)
	default:
		// unreachable
	}
}

func (parser *Parser) parse_number() {
	var value, err = strconv.ParseFloat(parser.prev_token.lexeme, 64)
	if err != nil {
		panic(err.Error())
	}
	parser.emit_constant(NumberValue(value))
}

func (parser *Parser) parse_string() {
	var end = len(parser.prev_token.lexeme) - 1
	var str = parser.prev_token.lexeme[1:end]
	var obj = Obj{spec: OBJ_STRING, data: []byte(str)}
	var val = ObjValue{ptr: &obj}
	parser.emit_constant(val)
}

func (parser *Parser) parse_grouping() {
	parser.parse_expression()
	parser.consume(TOKEN_RIGHT_PAREN, "expect ')' after expression")
}

func (parser *Parser) parse_unary() {

	var operator_spec = parser.prev_token.spec
	parser.parse_precedence(PREC_UNARY)

	switch operator_spec {
	case TOKEN_BANG:
		parser.emit_byte(OP_NOT)
	case TOKEN_MINUS:
		parser.emit_byte(OP_NEGATE)
	default:
		panic("unexpected operator spec")
	}
}

func (parser *Parser) parse_binary() {

	var operator_spec = parser.prev_token.spec

	var _, _, prec = get_rule(parser, operator_spec)
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

	case TOKEN_BANG_EQUAL:
		parser.emit_byte(OP_EQUAL, OP_NOT)
	case TOKEN_EQUAL_EQUAL:
		parser.emit_byte(OP_EQUAL)

	case TOKEN_GREATER:
		parser.emit_byte(OP_GREATER)
	case TOKEN_GREATER_EQUAL:
		parser.emit_byte(OP_LESS, OP_NOT)

	case TOKEN_LESS:
		parser.emit_byte(OP_LESS)
	case TOKEN_LESS_EQUAL:
		parser.emit_byte(OP_GREATER, OP_NOT)

	default:
		panic("unexpected operator spec")
	}

}

func (parser *Parser) parse_precedence(prec int) {

	parser.step()

	var prefix_rule, _, _ = get_rule(parser, parser.prev_token.spec)
	if prefix_rule == nil {
		parser.report_error("expect expression")
		return
	}
	prefix_rule()

	for {
		_, _, curr_prec := get_rule(parser, parser.curr_token.spec)
		if prec > curr_prec {
			break
		}

		parser.step()
		var _, infix_rule, _ = get_rule(parser, parser.prev_token.spec)
		infix_rule()
	}
}

func (parser *Parser) emit_byte(bs ...byte) {
	for _, b := range bs {
		var curr_chunk = parser.current_chunk()
		*curr_chunk = *write(curr_chunk, b, parser.prev_token.line)
	}
}

func (parser *Parser) emit_return() {
	parser.emit_byte(OP_RETURN)
}

func (parser *Parser) emit_constant(value Value) {
	parser.emit_byte(OP_CONSTANT, parser.make_constant(value))
}

func (parser *Parser) make_constant(value Value) byte {

	var curr_chunk = parser.current_chunk()

	var new_chunk, addr = add_constant(curr_chunk, value)

	if addr > 0xFF {
		parser.report_error("too many constants in one chunk")
		return 0x00
	}

	*curr_chunk = *new_chunk

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
