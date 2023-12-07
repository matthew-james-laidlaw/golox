package lox

import "fmt"

type Interpreter struct {
	Env     *Environment
	Globals *Environment
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		Env:     NewEnvironment(nil),
		Globals: NewEnvironment(nil),
	}
}

func (i Interpreter) VisitAssignment(expr *Assignment) interface{} {
	value := expr.Value.Accept(i)
	i.Env.Assign(expr.Name, value)
	return value
}

func (i Interpreter) VisitBinary(expr *Binary) interface{} {
	switch expr.Operation.Lexeme {
	case "+":
		return expr.Left.Accept(i).(float32) + expr.Right.Accept(i).(float32)
	case "-":
		return expr.Left.Accept(i).(float32) - expr.Right.Accept(i).(float32)
	case "*":
		return expr.Left.Accept(i).(float32) * expr.Right.Accept(i).(float32)
	case "/":
		return expr.Left.Accept(i).(float32) / expr.Right.Accept(i).(float32)
	case "==":
		return expr.Left.Accept(i) == expr.Right.Accept(i)
	case "!=":
		return expr.Left.Accept(i) != expr.Right.Accept(i)
	case ">":
		return expr.Left.Accept(i).(float32) > expr.Right.Accept(i).(float32)
	case ">=":
		return expr.Left.Accept(i).(float32) >= expr.Right.Accept(i).(float32)
	case "<":
		return expr.Left.Accept(i).(float32) < expr.Right.Accept(i).(float32)
	case "<=":
		return expr.Left.Accept(i).(float32) <= expr.Right.Accept(i).(float32)
	}
	return nil
}

func (i Interpreter) VisitCall(expr *Call) interface{} {
	callee := expr.Callee.Accept(i)

	arguments := make([]interface{}, 0)
	for _, argument := range expr.Arguments {
		arguments = append(arguments, argument.Accept(i))
	}

	function := callee.(Callable)
	return function.Call(i, arguments)
}

func (i Interpreter) VisitGet(expr *Get) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitGrouping(expr *Grouping) interface{} {
	return expr.Expr.Accept(i)
}

func (i Interpreter) VisitLiteral(expr *Literal) interface{} {
	return expr.Value
}

func (i Interpreter) VisitLogical(expr *Logical) interface{} {
	left := expr.Left.Accept(i)

	if expr.Operation.Type == OR {
		if IsTruthy(left) {
			return left
		}
	} else {
		if !IsTruthy(left) {
			return left
		}
	}

	return expr.Right.Accept(i)
}

func (i Interpreter) VisitSet(expr *Set) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitSuper(expr *Super) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitThis(expr *This) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitUnary(expr *Unary) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitVariable(expr *Variable) interface{} {
	return i.Env.Get(expr.Name)
}

func (i Interpreter) VisitBlock(stmt *Block) interface{} {
	i.ExecuteBlock(stmt.Statements, NewEnvironment(i.Env))
	return nil
}

func (i Interpreter) ExecuteBlock(statements []Statement, environment *Environment) {
	previous := i.Env

	i.Env = environment
	for _, statement := range statements {
		statement.Accept(i)
	}

	i.Env = previous
}

func (i Interpreter) VisitClass(stmt *Class) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitExpressionStatement(stmt *ExpressionStatement) interface{} {
	stmt.Expr.Accept(i)
	return nil
}

func (i Interpreter) VisitFunction(stmt *Function) interface{} {
	function := &LoxFunction{stmt}
	i.Env.Define(stmt.Name.Lexeme, function)
	return nil
}

func (i Interpreter) VisitIf(stmt *If) interface{} {
	if IsTruthy(stmt.Condition.Accept(i)) {
		stmt.ThenBranch.Accept(i)
	} else if stmt.ElseBranch != nil {
		stmt.ElseBranch.Accept(i)
	}
	return nil
}

func (i Interpreter) VisitPrint(stmt *Print) interface{} {
	fmt.Println(stmt.Expr.Accept(i))
	return nil
}

func (i Interpreter) VisitReturn(stmt *Return) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitVar(stmt *Var) interface{} {
	var value interface{}
	if stmt.Initializer != nil {
		value = stmt.Initializer.Accept(i)
	}
	i.Env.Define(stmt.Name.Lexeme, value)
	return nil
}

func (i Interpreter) VisitWhile(stmt *While) interface{} {
	for IsTruthy(stmt.Condition.Accept(i)) {
		stmt.Body.Accept(i)
	}
	return nil
}
