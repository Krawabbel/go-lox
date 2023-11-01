package lox

import "fmt"

var DEBUG_TRACE_EXECUTION = true
var DEBUG_PRINT_CODE = true

func disassemble_chunk(chunk *Chunk, name string) string {
	prog := fmt.Sprintf("== %s ==\n", name)

	for addr := 0; addr < count(chunk); {
		dasm, offset := disassemble_instruction(chunk, addr)
		addr += offset
		prog += dasm + "\n"
	}
	return prog
}

func disassemble_instruction(chunk *Chunk, addr int) (string, int) {
	info := fmt.Sprintf("%04d ", addr)

	if addr > 0 && chunk.lines[addr] == chunk.lines[addr-1] {
		info += "   | "
	} else {
		info += fmt.Sprintf("%4d ", chunk.lines[addr])
	}

	instr := chunk.code[addr]
	switch instr {

	case OP_ADD:
		var dasm, offset = disassemble_simple_instruction("OP_ADD")
		return info + dasm, offset

	case OP_SUBTRACT:
		var dasm, offset = disassemble_simple_instruction("OP_SUBTRACT")
		return info + dasm, offset

	case OP_MULTIPLY:
		var dasm, offset = disassemble_simple_instruction("OP_MULTIPLY")
		return info + dasm, offset

	case OP_DIVIDE:
		var dasm, offset = disassemble_simple_instruction("OP_DIVIDE")
		return info + dasm, offset

	case OP_NOT:
		var dasm, offset = disassemble_simple_instruction("OP_NOT")
		return info + dasm, offset

	case OP_NEGATE:
		var dasm, offset = disassemble_simple_instruction("OP_NEGATE")
		return info + dasm, offset

	case OP_RETURN:
		var dasm, offset = disassemble_simple_instruction("OP_RETURN")
		return info + dasm, offset

	case OP_CONSTANT:
		var dasm, offset = disassemble_constant_instruction("OP_CONSTANT", chunk, addr+1)
		return info + dasm, offset

	case OP_NIL:
		var dasm, offset = disassemble_simple_instruction("OP_NIL")
		return info + dasm, offset

	case OP_TRUE:
		var dasm, offset = disassemble_simple_instruction("OP_TRUE")
		return info + dasm, offset

	case OP_FALSE:
		var dasm, offset = disassemble_simple_instruction("OP_FALSE")
		return info + dasm, offset

	case OP_EQUAL:
		var dasm, offset = disassemble_simple_instruction("OP_EQUAL")
		return info + dasm, offset

	case OP_GREATER:
		var dasm, offset = disassemble_simple_instruction("OP_GREATER")
		return info + dasm, offset

	case OP_LESS:
		var dasm, offset = disassemble_simple_instruction("OP_LESS")
		return info + dasm, offset
	}
	panic("unknown opcode")
}

func disassemble_simple_instruction(name string) (string, int) {
	return name, 1
}

func disassemble_constant_instruction(name string, chunk *Chunk, code_addr int) (string, int) {
	var addr = chunk.code[code_addr]
	var val = chunk.constants[addr]

	return fmt.Sprintf("%-16s (%04d) %s", name, addr, val.stringify()), 2
}
