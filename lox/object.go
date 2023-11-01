package lox

const (
	OBJ_STRING = iota
)

type Obj struct {
	spec int
	data []byte
}

func is_string(val Value) bool {
	return is_obj(val) && ((val.(ObjValue)).ptr.spec == OBJ_STRING)
}

func (obj Obj) stringify() string {
	switch obj.spec {
	case OBJ_STRING:
		return string(obj.data)
	}
	panic("unexpected object spec")
}
