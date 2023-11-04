package lox

const (
	OP_RETURN = iota
	OP_CONSTANT
	OP_NIL
	OP_TRUE
	OP_FALSE
	OP_POP
	OP_DEFINE_GLOBAL
	OP_EQUAL
	OP_GREATER
	OP_LESS
	OP_ADD
	OP_SUBTRACT
	OP_MULTIPLY
	OP_DIVIDE
	OP_NOT
	OP_NEGATE
	OP_PRINT
)

type Chunk struct {
	code      []byte
	lines     []int
	constants []Value
}

func write(chunk *Chunk, b byte, line int) *Chunk {
	if chunk.code == nil {
		chunk.code = make([]byte, 0, 1)
		chunk.lines = make([]int, 0, 1)
	}
	chunk.code = append(chunk.code, b)
	chunk.lines = append(chunk.lines, line)
	return chunk
}

func add_constant(chunk *Chunk, val Value) (*Chunk, int) {
	if chunk.constants == nil {
		chunk.constants = make([]Value, 0, 1)
	}
	chunk.constants = append(chunk.constants, val)
	return chunk, len(chunk.constants) - 1
}

func count(chunk *Chunk) int {
	return len(chunk.code)
}
