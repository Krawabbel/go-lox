package lox

import "fmt"

const (
	INTERPRET_OK = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

type VM struct {
	chunk *Chunk
	stack Stack[Value]
	ip    int
}

func (vm *VM) interpret(src string) int {

	var chunk, success = compile(src)
	if !success {
		return INTERPRET_COMPILE_ERROR
	}

	vm.chunk = chunk
	vm.ip = 0

	return vm.run()
}

func (vm *VM) next_code() byte {
	b := vm.chunk.code[vm.ip]
	vm.ip++
	return b
}

func (vm *VM) next_constant() Value {
	return vm.chunk.constants.values[vm.next_code()]
}

func (vm *VM) execute_binary(f func(a, b Value) Value) error {

	b, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	vm.stack.Push(f(a, b))
	return nil

}

func (vm *VM) run() int {
	for {

		if DEBUG_TRACE_EXECUTION {
			dbg, _ := disassemble_instruction(*vm.chunk, vm.ip)
			debug(fmt.Sprintf("%-50s %s", dbg, vm.stack.Dump()))
		}

		instr := vm.next_code()

		switch instr {

		case OP_ADD:
			if err := vm.execute_binary(func(a, b Value) Value { return a + b }); err != nil {
				return INTERPRET_RUNTIME_ERROR
			}

		case OP_SUBTRACT:
			if err := vm.execute_binary(func(a, b Value) Value { return a - b }); err != nil {
				return INTERPRET_RUNTIME_ERROR
			}

		case OP_MULTIPLY:
			if err := vm.execute_binary(func(a, b Value) Value { return a * b }); err != nil {
				return INTERPRET_RUNTIME_ERROR
			}

		case OP_DIVIDE:
			if err := vm.execute_binary(func(a, b Value) Value { return a / b }); err != nil {
				return INTERPRET_RUNTIME_ERROR
			}

		case OP_NEGATE:
			val, err := vm.stack.Pop()
			if err != nil {
				return INTERPRET_RUNTIME_ERROR
			}
			vm.stack.Push(-val)

		case OP_CONSTANT:
			constant := vm.next_constant()
			vm.stack.Push(constant)

		case OP_RETURN:
			addr, err := vm.stack.Pop()
			if err != nil {
				return INTERPRET_RUNTIME_ERROR
			}
			print("%s\n", stringify(addr))
			return INTERPRET_OK

		default:
			return INTERPRET_RUNTIME_ERROR
		}
	}
}
