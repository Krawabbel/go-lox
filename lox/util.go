package lox

func if_then_else[T any](b bool, t_true, t_false T) T {
	if b {
		return t_true
	}
	return t_false
}
