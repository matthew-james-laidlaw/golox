package statement

type Visitor interface {
	VisitBlock(stmt *Block) interface{}
	VisitClass(stmt *Class) interface{}
	VisitExpression(stmt *Expression) interface{}
	VisitFunction(stmt *Function) interface{}
	VisitIf(stmt *If) interface{}
	VisitPrint(stmt *Print) interface{}
	VisitReturn(stmt *Return) interface{}
	VisitVar(stmt *Var) interface{}
	VisitWhile(stmt *While) interface{}
}
