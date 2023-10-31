package lox

import "fmt"

type Value = float64

func stringify(val Value) string {
	return fmt.Sprintf("%g", val)
}

type value_array struct {
	values []Value
}

func (va value_array) count() int {
	return len(va.values)
}

func (va *value_array) write(b Value) {
	if va.values == nil {
		va.values = make([]Value, 0)
	}
	va.values = append(va.values, b)
}
