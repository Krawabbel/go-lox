package lox

const (
	OP_RETURN = iota
	OP_CONSTANT
	OP_NEGATE
	OP_ADD
	OP_SUBTRACT
	OP_MULTIPLY
	OP_DIVIDE
)

type Chunk struct {
	code      []byte
	lines     []int
	constants value_array
}

func (chunk Chunk) count() int {
	return len(chunk.code)
}

func (chunk *Chunk) write(b byte, line int) {
	if chunk.code == nil {
		chunk.code = make([]byte, 0)
		chunk.lines = make([]int, 0)
	}
	chunk.code = append(chunk.code, b)
	chunk.lines = append(chunk.lines, line)
}

func (chunk *Chunk) add_constant(val Value) int {
	chunk.constants.write(val)
	return chunk.constants.count() - 1
}
