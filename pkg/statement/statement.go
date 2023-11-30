package statement

type Statement interface {
	Accept(v Visitor) interface{}
}

type Block struct{}

func (stmt *Block) Accept(v Visitor) interface{} {
	return v.VisitBlock(stmt)
}

type Class struct{}

func (stmt *Class) Accept(v Visitor) interface{} {
	return v.VisitClass(stmt)
}

type Expression struct{}

func (stmt *Expression) Accept(v Visitor) interface{} {
	return v.VisitExpression(stmt)
}

type Function struct{}

func (stmt *Function) Accept(v Visitor) interface{} {
	return v.VisitFunction(stmt)
}

type If struct{}

func (stmt *If) Accept(v Visitor) interface{} {
	return v.VisitIf(stmt)
}

type Print struct{}

func (stmt *Print) Accept(v Visitor) interface{} {
	return v.VisitPrint(stmt)
}

type Return struct{}

func (stmt *Return) Accept(v Visitor) interface{} {
	return v.VisitReturn(stmt)
}

type Var struct{}

func (stmt *Var) Accept(v Visitor) interface{} {
	return v.VisitVar(stmt)
}

type While struct{}

func (stmt *While) Accept(v Visitor) interface{} {
	return v.VisitWhile(stmt)
}
