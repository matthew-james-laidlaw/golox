package expression

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
}
