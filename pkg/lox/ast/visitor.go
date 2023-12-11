package ast

type Visitor interface {
	VisitAssignment(expr *Assignment) interface{}
	VisitBinary(expr *Binary) interface{}
	VisitCall(expr *Call) interface{}
	VisitGet(expr *Get) interface{}
	VisitGrouping(expr *Grouping) interface{}
	VisitLiteral(expr *Literal) interface{}
	VisitLogical(expr *Logical) interface{}
	VisitSet(expr *Set) interface{}
	VisitSuper(expr *Super) interface{}
	VisitThis(expr *This) interface{}
	VisitUnary(expr *Unary) interface{}
	VisitVariable(expr *Variable) interface{}

	VisitBlock(stmt *Block) interface{}
	VisitClass(stmt *Class) interface{}
	VisitExpressionStatement(stmt *ExpressionStatement) interface{}
	VisitFunction(stmt *Function) interface{}
	VisitIf(stmt *If) interface{}
	VisitPrint(stmt *Print) interface{}
	VisitReturn(stmt *Return) interface{}
	VisitVar(stmt *Var) interface{}
	VisitWhile(stmt *While) interface{}
}
