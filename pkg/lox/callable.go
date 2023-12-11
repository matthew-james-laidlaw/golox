package lox

import "golox/pkg/lox/ast"

type Callable interface {
	Call(i Interpreter, arguments []interface{}) interface{}
}

type LoxFunction struct {
	Declaration *ast.Function
}

func (f *LoxFunction) Call(i Interpreter, arguments []interface{}) (returnValue interface{}) {
	environment := NewEnvironment(i.Globals)
	for i, _ := range f.Declaration.Params {
		environment.Define(f.Declaration.Params[i].Lexeme, arguments[i])
	}

	defer func() {
		err := recover()

		if err == nil {
			return
		}

		v, ok := err.(ReturnValue)
		if !ok {
			panic(err)
		}

		returnValue = v.Value
		return
	}()

	i.ExecuteBlock(f.Declaration.Body, environment)

	return nil
}
