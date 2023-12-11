package ast

import (
	"golox/pkg/lox/token"
)

type Statement interface {
	Accept(v Visitor) interface{}
}

type Block struct {
	Statements []Statement
}

func (stmt Block) Accept(v Visitor) interface{} {
	return v.VisitBlock(&stmt)
}

type Class struct {
	Superclass Variable
	Methods    []Function
}

func (stmt Class) Accept(v Visitor) interface{} {
	return v.VisitClass(&stmt)
}

type ExpressionStatement struct {
	Expr Expression
}

func (stmt ExpressionStatement) Accept(v Visitor) interface{} {
	return v.VisitExpressionStatement(&stmt)
}

type Function struct {
	Name   token.Token
	Params []token.Token
	Body   []Statement
}

func (stmt Function) Accept(v Visitor) interface{} {
	return v.VisitFunction(&stmt)
}

type If struct {
	Condition  Expression
	ThenBranch Statement
	ElseBranch Statement
}

func (stmt If) Accept(v Visitor) interface{} {
	return v.VisitIf(&stmt)
}

type Print struct {
	Expr Expression
}

func (stmt Print) Accept(v Visitor) interface{} {
	return v.VisitPrint(&stmt)
}

type Return struct {
	Value Expression
}

func (stmt Return) Accept(v Visitor) interface{} {
	return v.VisitReturn(&stmt)
}

type Var struct {
	Name        token.Token
	Initializer Expression
}

func (stmt Var) Accept(v Visitor) interface{} {
	return v.VisitVar(&stmt)
}

type While struct {
	Condition Expression
	Body      Statement
}

func (stmt While) Accept(v Visitor) interface{} {
	return v.VisitWhile(&stmt)
}
