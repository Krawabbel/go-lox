package lox

import "fmt"

const (
	VAL_NIL = iota
	VAL_BOOL
	VAL_NUMBER
)

type Value interface {
	as_string() string
	as_number() NumberValue
	as_bool() bool

	get_type() int
}

func has_type(val Value, t int) bool {
	return val.get_type() == t
}

func is_falsey(val Value) bool {
	return has_type(val, VAL_NIL) || (has_type(val, VAL_BOOL) && !val.as_bool())
}

func are_equal(val_1, val_2 Value) bool {
	if val_1.get_type() != val_2.get_type() {
		return false
	}

	switch val_1.get_type() {

	case VAL_BOOL:
		return val_1.as_bool() == val_2.as_bool()

	case VAL_NIL:
		return true

	case VAL_NUMBER:
		return val_1.as_number() == val_2.as_number()
	}

	return false // unreachable
}

type NilValue struct{}

func (val NilValue) as_string() string {
	return "nil"
}

func (val NilValue) as_number() NumberValue {
	return 0
}

func (val NilValue) as_bool() bool {
	return false
}

func (val NilValue) get_type() int {
	return VAL_NIL
}

type BoolValue bool

func (val BoolValue) as_string() string {
	return fmt.Sprintf("%t", val)
}

func (val BoolValue) as_number() NumberValue {
	return if_then_else[NumberValue](bool(val), 1, 0)
}

func (val BoolValue) as_bool() bool {
	return bool(val)
}

func (val BoolValue) get_type() int {
	return VAL_BOOL
}

type NumberValue float64

func (val NumberValue) as_string() string {
	return fmt.Sprintf("%g", val)
}

func (val NumberValue) as_number() NumberValue {
	return val
}

func (val NumberValue) as_bool() bool {
	return val != 0
}

func (val NumberValue) get_type() int {
	return VAL_NUMBER
}
