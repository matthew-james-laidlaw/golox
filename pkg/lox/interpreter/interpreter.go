package interpreter

import (
	"fmt"
	"golox/pkg/lox/ast"
	"golox/pkg/lox/token"
)

type Interpreter struct {
	Env     *Environment
	Globals *Environment
}

func NewInterpreter() *Interpreter {
	env := NewEnvironment(nil)
	return &Interpreter{
		Env:     env,
		Globals: env,
	}
}

func (i Interpreter) VisitAssignment(expr *ast.Assignment) interface{} {
	value := expr.Value.Accept(i)
	i.Env.Assign(expr.Name, value)
	return value
}

func (i Interpreter) VisitBinary(expr *ast.Binary) interface{} {
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

func (i Interpreter) VisitCall(expr *ast.Call) interface{} {
	callee := expr.Callee.Accept(i)

	arguments := make([]interface{}, 0)
	for _, argument := range expr.Arguments {
		arguments = append(arguments, argument.Accept(i))
	}

	function := callee.(Callable)
	return function.Call(i, arguments)
}

func (i Interpreter) VisitGet(expr *ast.Get) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitGrouping(expr *ast.Grouping) interface{} {
	return expr.Expr.Accept(i)
}

func (i Interpreter) VisitLiteral(expr *ast.Literal) interface{} {
	return expr.Value
}

func (i Interpreter) VisitLogical(expr *ast.Logical) interface{} {
	left := expr.Left.Accept(i)

	if expr.Operation.Type == token.OR {
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

func (i Interpreter) VisitSet(expr *ast.Set) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitSuper(expr *ast.Super) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitThis(expr *ast.This) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitUnary(expr *ast.Unary) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitVariable(expr *ast.Variable) interface{} {
	return i.Env.Get(expr.Name)
}

func (i Interpreter) VisitBlock(stmt *ast.Block) interface{} {
	i.ExecuteBlock(stmt.Statements, NewEnvironment(i.Env))
	return nil
}

func (i Interpreter) ExecuteBlock(statements []ast.Statement, environment *Environment) {
	previous := i.Env

	i.Env = environment
	for _, statement := range statements {
		statement.Accept(i)
	}

	i.Env = previous
}

func (i Interpreter) VisitClass(stmt *ast.Class) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitExpressionStatement(stmt *ast.ExpressionStatement) interface{} {
	stmt.Expr.Accept(i)
	return nil
}

func (i Interpreter) VisitFunction(stmt *ast.Function) interface{} {
	function := &LoxFunction{stmt}
	i.Env.Define(stmt.Name.Lexeme, function)
	return nil
}

func (i Interpreter) VisitIf(stmt *ast.If) interface{} {
	if IsTruthy(stmt.Condition.Accept(i)) {
		stmt.ThenBranch.Accept(i)
	} else if stmt.ElseBranch != nil {
		stmt.ElseBranch.Accept(i)
	}
	return nil
}

func (i Interpreter) VisitPrint(stmt *ast.Print) interface{} {
	fmt.Println(stmt.Expr.Accept(i))
	return nil
}

type ReturnValue struct {
	Value interface{}
}

func (i Interpreter) VisitReturn(stmt *ast.Return) interface{} {
	var value interface{}
	if stmt.Value != nil {
		value = stmt.Value.Accept(i)
	}
	panic(ReturnValue{value})
}

func (i Interpreter) VisitVar(stmt *ast.Var) interface{} {
	var value interface{}
	if stmt.Initializer != nil {
		value = stmt.Initializer.Accept(i)
	}
	i.Env.Define(stmt.Name.Lexeme, value)
	return nil
}

func (i Interpreter) VisitWhile(stmt *ast.While) interface{} {
	for IsTruthy(stmt.Condition.Accept(i)) {
		stmt.Body.Accept(i)
	}
	return nil
}
