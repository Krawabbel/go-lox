package lox

import "fmt"

const (
	INTERPRET_OK = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

type VM struct {
	chunk *Chunk
	stack Stack
	ip    int

	globals map[string]Value
}

func newVM() *VM {
	var vm = new(VM)
	vm.globals = make(map[string]Value)
	return vm
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
	vm.ip++
	return vm.chunk.code[vm.ip-1]
}

func (vm *VM) next_constant() Value {
	return vm.chunk.constants[vm.next_code()]
}

func (vm *VM) next_string() string {
	var constant = vm.next_constant()
	return constant.(ObjValue).ptr.stringify()
}

func (vm *VM) run() int {
	for {

		if DEBUG_TRACE_EXECUTION {
			dbg, _ := disassemble_instruction(vm.chunk, vm.ip)
			fmt.Fprintf(STDDBG, "%-50s %s\n", dbg, vm.stack.dump())
		}

		instr := vm.next_code()

		switch instr {

		case OP_ADD:

			switch {

			case is_string(vm.stack.peek(0)) && is_string(vm.stack.peek(1)):
				var b = string(vm.stack.pop().(ObjValue).ptr.data)
				var a = string(vm.stack.pop().(ObjValue).ptr.data)
				var obj = Obj{spec: OBJ_STRING, data: []byte(a + b)}
				var val = ObjValue{ptr: &obj}
				vm.stack.push(val)

			case is_number(vm.stack.peek(0)) && is_number(vm.stack.peek(1)):
				var b = vm.stack.pop().(NumberValue)
				var a = vm.stack.pop().(NumberValue)
				vm.stack.push(a + b)

			default:
				vm.report_runtime_error("operands must be two numbers or two strings")
				return INTERPRET_RUNTIME_ERROR
			}

		case OP_SUBTRACT:
			if !vm.execute_binary(func(a, b NumberValue) Value { return a - b }) {
				return INTERPRET_RUNTIME_ERROR
			}

		case OP_MULTIPLY:
			if !vm.execute_binary(func(a, b NumberValue) Value { return a * b }) {
				return INTERPRET_RUNTIME_ERROR
			}

		case OP_DIVIDE:
			if !vm.execute_binary(func(a, b NumberValue) Value { return a / b }) {
				return INTERPRET_RUNTIME_ERROR
			}

		case OP_NOT:
			vm.stack.push(BoolValue(is_falsey(vm.stack.pop())))

		case OP_NEGATE:
			if !is_number(vm.stack.peek(0)) {
				vm.report_runtime_error("operand must be a number")
				return INTERPRET_RUNTIME_ERROR
			}
			vm.stack.push(-vm.stack.pop().(NumberValue))

		case OP_CONSTANT:
			var constant = vm.next_constant()
			vm.stack.push(constant)

		case OP_NIL:
			vm.stack.push(NilValue{})
		case OP_TRUE:
			vm.stack.push(BoolValue(true))
		case OP_FALSE:
			vm.stack.push(BoolValue(false))

		case OP_POP:
			_ = vm.stack.pop()

		case OP_DEFINE_GLOBAL:
			var name = vm.next_string()
			vm.globals[name] = vm.stack.peek(0)
			_ = vm.stack.pop()

		case OP_EQUAL:
			var b = vm.stack.pop()
			var a = vm.stack.pop()
			vm.stack.push(BoolValue(are_equal(a, b)))

		case OP_GREATER:
			vm.execute_binary(func(a, b NumberValue) Value { return BoolValue(a > b) })

		case OP_LESS:
			vm.execute_binary(func(a, b NumberValue) Value { return BoolValue(a < b) })

		case OP_PRINT:
			fmt.Fprintf(STDOUT, "%s\n", vm.stack.pop().stringify())

		case OP_RETURN:
			// exit interpreter
			return INTERPRET_OK

		default:
			panic("unexpected opcode")
		}
	}
}

func (vm *VM) report_runtime_error(format string, a ...any) {
	fmt.Fprintf(STDERR, format, a...)

	var instr = vm.ip - 1
	var line = vm.chunk.lines[instr]

	fmt.Fprintf(STDERR, "\n[line %d] in script\n", line)

	vm.stack.reset()
}

func (vm *VM) execute_binary(f func(a, b NumberValue) Value) (success bool) {

	if !is_number(vm.stack.peek(0)) || !is_number(vm.stack.peek(1)) {
		vm.report_runtime_error("operands must be two numbers")
		return false
	}

	var b = vm.stack.pop().(NumberValue)
	var a = vm.stack.pop().(NumberValue)

	var res = f(a, b)

	vm.stack.push(res)

	return true
}
