package lox

import (
	"strings"
)

type Stack struct {
	data []Value
	ptr  int
}

func (stack *Stack) push(val Value) {

	if stack.data == nil {
		stack.data = make([]Value, 0, 1)
	}

	if stack.ptr == len(stack.data) {
		stack.data = append(stack.data, val)
	} else {
		stack.data[stack.ptr] = val
	}

	stack.ptr++
}

func (stack *Stack) pop() Value {
	stack.ptr--
	return stack.data[stack.ptr]
}

func (stack *Stack) reset() {
	stack.ptr = 0
}

func (stack Stack) peek(pos int) Value {
	return stack.data[stack.ptr-1-pos]
}

func (stack Stack) dump() string {
	var dumps = make([]string, stack.ptr)
	for i := range dumps {
		dumps[len(dumps)-1-i] = stack.peek(i).stringify()
	}
	return "[" + strings.Join(dumps, ", ") + "]"
}
