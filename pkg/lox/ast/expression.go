package ast

import (
	"golox/pkg/lox/token"
)

type Expression interface {
	Accept(v Visitor) interface{}
}

type Assignment struct {
	Name  token.Token
	Value Expression
}

func (expr Assignment) Accept(v Visitor) interface{} {
	return v.VisitAssignment(&expr)
}

type Binary struct {
	Operation token.Token
	Left      Expression
	Right     Expression
}

func (expr Binary) Accept(v Visitor) interface{} {
	return v.VisitBinary(&expr)
}

type Call struct {
	Callee    Expression
	Arguments []Expression
}

func (expr Call) Accept(v Visitor) interface{} {
	return v.VisitCall(&expr)
}

type Get struct{}

func (expr Get) Accept(v Visitor) interface{} {
	return v.VisitGet(&expr)
}

type Grouping struct {
	Expr Expression
}

func (expr Grouping) Accept(v Visitor) interface{} {
	return v.VisitGrouping(&expr)
}

type Literal struct {
	Value interface{}
}

func (expr Literal) Accept(v Visitor) interface{} {
	return v.VisitLiteral(&expr)
}

type Logical struct {
	Operation token.Token
	Left      Expression
	Right     Expression
}

func (expr Logical) Accept(v Visitor) interface{} {
	return v.VisitLogical(&expr)
}

type Set struct{}

func (expr Set) Accept(v Visitor) interface{} {
	return v.VisitSet(&expr)
}

type Super struct{}

func (expr Super) Accept(v Visitor) interface{} {
	return v.VisitSuper(&expr)
}

type This struct{}

func (expr This) Accept(v Visitor) interface{} {
	return v.VisitThis(&expr)
}

type Unary struct {
	Operation token.Token
	Operand   Expression
}

func (expr Unary) Accept(v Visitor) interface{} {
	return v.VisitUnary(&expr)
}

type Variable struct {
	Name token.Token
}

func (expr Variable) Accept(v Visitor) interface{} {
	return v.VisitVariable(&expr)
}
