package lox

type Environment struct {
	Values    map[string]interface{}
	Enclosing *Environment
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		Values:    make(map[string]interface{}),
		Enclosing: enclosing,
	}
}

func (e *Environment) Assign(name Token, value interface{}) {
	_, ok := e.Values[name.Lexeme]

	if ok {
		e.Values[name.Lexeme] = value
	} else if e.Enclosing != nil {
		e.Enclosing.Assign(name, value)
	} else {
		panic("undefined variable")
	}
}

func (e *Environment) Define(name string, value interface{}) {
	e.Values[name] = value
}

func (e *Environment) Get(name Token) interface{} {
	value, ok := e.Values[name.Lexeme]
	if ok {
		return value
	} else if e.Enclosing != nil {
		return e.Enclosing.Get(name)
	} else {
		panic("undefined variable")
	}
}
