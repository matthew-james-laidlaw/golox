package interpreter

import (
	"golox/pkg/expression"
	"golox/pkg/statement"
)

type Interpreter struct {
	Variables map[string]interface{}
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		Variables: make(map[string]interface{}),
	}
}

func (i Interpreter) VisitAssignment(expr *expression.Assignment) interface{} {
	i.Variables[expr.Name.Lexeme] = expr.Value.Accept(i)
	return nil
}

func (i Interpreter) VisitBinary(expr *expression.Binary) interface{} {
	switch expr.Operation.Lexeme {
	case "+":
		return expr.Left.Accept(i).(float32) + expr.Right.Accept(i).(float32)
	case "-":
		return expr.Left.Accept(i).(float32) - expr.Right.Accept(i).(float32)
	case "*":
		return expr.Left.Accept(i).(float32) * expr.Right.Accept(i).(float32)
	case "/":
		return expr.Left.Accept(i).(float32) / expr.Right.Accept(i).(float32)
	}
	return nil
}

func (i Interpreter) VisitCall(expr *expression.Call) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitGet(expr *expression.Get) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitGrouping(expr *expression.Grouping) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitLiteral(expr *expression.Literal) interface{} {
	return expr.Value
}

func (i Interpreter) VisitLogical(expr *expression.Logical) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitSet(expr *expression.Set) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitSuper(expr *expression.Super) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitThis(expr *expression.This) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitUnary(expr *expression.Unary) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitVariable(expr *expression.Variable) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitBlock(stmt *statement.Block) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitClass(stmt *statement.Class) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitExpression(stmt *statement.Expression) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitFunction(stmt *statement.Function) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitIf(stmt *statement.If) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitPrint(stmt *statement.Print) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitReturn(stmt *statement.Return) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitVar(stmt *statement.Var) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitWhile(stmt *statement.While) interface{} {
	//TODO implement me
	panic("implement me")
}
