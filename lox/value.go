package lox

import (
	"fmt"
)

const (
	VAL_NIL = iota
	VAL_BOOL
	VAL_NUMBER
	VAL_OBJ
	VAL_UNKNOWN
)

type Value interface {
	stringify() string
}

func get_type(val Value) int {
	switch val.(type) {
	case NilValue:
		return VAL_NIL
	case BoolValue:
		return VAL_BOOL
	case NumberValue:
		return VAL_NUMBER
	case ObjValue:
		return VAL_OBJ
	}
	panic("unexpected value type")
}

func is_falsey(val Value) bool {
	return is_nil(val) || (is_bool(val) && !bool(val.(BoolValue)))
}

func are_equal(val_1, val_2 Value) bool {
	if get_type(val_1) != get_type(val_2) {
		return false
	}

	switch val_1.(type) {

	case BoolValue:
		return val_1.(BoolValue) == val_2.(BoolValue)

	case NilValue:
		return true

	case NumberValue:
		return val_1.(NumberValue) == val_2.(NumberValue)

	case ObjValue:
		var a = string(val_1.(ObjValue).ptr.data)
		var b = string(val_2.(ObjValue).ptr.data)
		return a == b
	}
	panic("unexpected value type")
}

func is_nil(val Value) bool {
	return get_type(val) == VAL_NIL
}

func is_number(val Value) bool {
	return get_type(val) == VAL_NUMBER
}

func is_bool(val Value) bool {
	return get_type(val) == VAL_BOOL
}

func is_obj(val Value) bool {
	return get_type(val) == VAL_OBJ
}

type NilValue struct{}

func (val NilValue) stringify() string {
	return "nil"
}

type BoolValue bool

func (val BoolValue) stringify() string {
	return fmt.Sprintf("%t", val)
}

type NumberValue float64

func (val NumberValue) stringify() string {
	return fmt.Sprintf("%g", val)
}

type ObjValue struct {
	ptr *Obj
}

func (val ObjValue) stringify() string {
	return val.ptr.stringify()
}
