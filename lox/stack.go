package lox

import "fmt"

type Stack[T any] struct {
	data []T
	ptr  int
}

func (stack *Stack[T]) Push(t T) {

	if stack.data == nil {
		stack.data = make([]T, 0)
	}

	if stack.ptr == len(stack.data) {
		stack.data = append(stack.data, t)
	} else {
		stack.data[stack.ptr] = t
	}

	stack.ptr++
}

func (stack *Stack[T]) Pop() (t T, err error) {
	if len(stack.data) == 0 {
		return t, fmt.Errorf("cannot pop from empty stack")
	}
	stack.ptr--
	return stack.data[stack.ptr], nil
}

func (stack *Stack[T]) Reset() {
	stack.ptr = 0
}

func (stack Stack[T]) Dump() string {
	return fmt.Sprintf("%v", stack.data[:stack.ptr])
}
