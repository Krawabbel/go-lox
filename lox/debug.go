package lox

import "fmt"

var DEBUG_TRACE_EXECUTION = true
var DEBUG_PRINT_CODE = true

func debug(msg string) {
	fmt.Println(msg)
}

func disassemble_chunk(chunk Chunk, name string) string {
	prog := fmt.Sprintf("== %s ==\n", name)

	for addr := 0; addr < chunk.count(); {
		dasm, offset := disassemble_instruction(chunk, addr)
		addr += offset
		prog += dasm + "\n"
	}
	return prog
}

func disassemble_instruction(chunk Chunk, addr int) (string, int) {
	info := fmt.Sprintf("%04d ", addr)

	if addr > 0 && chunk.lines[addr] == chunk.lines[addr-1] {
		info += "   | "
	} else {
		info += fmt.Sprintf("%4d ", chunk.lines[addr])
	}

	instr := chunk.code[addr]
	switch instr {

	case OP_ADD:
		dasm, offset := simple_instruction("OP_ADD")
		return info + dasm, offset

	case OP_SUBTRACT:
		dasm, offset := simple_instruction("OP_SUBTRACT")
		return info + dasm, offset

	case OP_MULTIPLY:
		dasm, offset := simple_instruction("OP_MULTIPLY")
		return info + dasm, offset

	case OP_DIVIDE:
		dasm, offset := simple_instruction("OP_DIVIDE")
		return info + dasm, offset

	case OP_NEGATE:
		dasm, offset := simple_instruction("OP_NEGATE")
		return info + dasm, offset

	case OP_RETURN:
		dasm, offset := simple_instruction("OP_RETURN")
		return info + dasm, offset

	case OP_CONSTANT:
		dasm, offset := constant_instruction("OP_CONSTANT", chunk, addr+1)
		return info + dasm, offset

	default:
		return info + fmt.Sprintf("unknown opcode %d", instr), 1
	}
}

func simple_instruction(name string) (string, int) {
	return name, 1
}

func constant_instruction(name string, chunk Chunk, code_addr int) (string, int) {
	addr := chunk.code[code_addr]
	val := chunk.constants.values[addr]

	return fmt.Sprintf("%-16s (%04d) %s", name, addr, stringify(val)), 2
}
