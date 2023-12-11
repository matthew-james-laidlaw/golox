package interpreter

type Boolean = bool
type Number = float32
type String = string

func IsTruthy(value interface{}) bool {
	switch value.(type) {
	case Boolean:
		return value.(Boolean) == true
	case Number:
		return value.(Number) != Number(0.0)
	case String:
		return len(value.(String)) != 0
	case nil:
		return false
	default:
		panic("invalid lox type")
	}
}
