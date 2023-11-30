package main

import (
	"fmt"
	"golox/pkg/expression"
	"golox/pkg/parser"
)

func main() {

	left := expression.Literal{Value: float32(40.0)}

	right := expression.Literal{Value: float32(2.0)}

	binary := expression.Binary{
		Operation: parser.Token{Lexeme: "+"},
		Left:      left,
		Right:     right,
	}

	interpreter := NewInterpreter()

	result := binary.Accept(interpreter)

	fmt.Println(result)

}
