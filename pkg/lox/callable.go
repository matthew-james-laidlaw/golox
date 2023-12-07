package lox

type Callable interface {
	Call(i Interpreter, arguments []interface{}) interface{}
}

type LoxFunction struct {
	Declaration *Function
}

func (f *LoxFunction) Call(i Interpreter, arguments []interface{}) interface{} {
	environment := NewEnvironment(i.Globals)
	for i, _ := range f.Declaration.Params {
		environment.Define(f.Declaration.Params[i].Lexeme, arguments[i])
	}
	i.ExecuteBlock(f.Declaration.Body, environment)
	return nil
}
