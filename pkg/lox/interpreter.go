package lox

import "fmt"

type Interpreter struct {
	Variables map[string]interface{}
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		Variables: make(map[string]interface{}),
	}
}

func (i Interpreter) VisitAssignment(expr *Assignment) interface{} {
	value := expr.Value.Accept(i)
	i.Variables[expr.Name.Lexeme] = value
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
	}
	return nil
}

func (i Interpreter) VisitCall(expr *Call) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitGet(expr *Get) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitGrouping(expr *Grouping) interface{} {
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitLiteral(expr *Literal) interface{} {
	return expr.Value
}

func (i Interpreter) VisitLogical(expr *Logical) interface{} {
	//TODO implement me
	panic("implement me")
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
	return i.Variables[expr.Name.Lexeme]
}

func (i Interpreter) VisitBlock(stmt *Block) interface{} {
	for _, statement := range stmt.Statements {
		statement.Accept(i)
	}
	return nil
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
	//TODO implement me
	panic("implement me")
}

func (i Interpreter) VisitIf(stmt *If) interface{} {
	//TODO implement me
	panic("implement me")
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
	i.Variables[stmt.Name.Lexeme] = value
	return nil
}

func (i Interpreter) VisitWhile(stmt *While) interface{} {
	for stmt.Condition.Accept(i).(bool) == true {
		stmt.Body.Accept(i)
	}
	return nil
}
