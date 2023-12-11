package main

import (
	"bufio"
	"fmt"
	"golox/pkg/lox/interpreter"
	"golox/pkg/lox/parser"
	"golox/pkg/lox/scanner"
	"log"
	"os"
)

type Lox struct {
	Interpreter *interpreter.Interpreter
}

func NewLox() *Lox {
	return &Lox{
		Interpreter: interpreter.NewInterpreter(),
	}
}

func (l *Lox) Main(args []string) {
	if len(args) > 2 {
		log.Fatalln("usage: golox [script]")
	} else if len(args) == 2 {
		l.RunFile(args[1])
	} else {
		l.RunPrompt()
	}
}

func (l *Lox) RunFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	l.Run(string(bytes))
}

func (l *Lox) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		l.Run(line)
	}
}

func (l *Lox) Run(source string) {
	tokens := scanner.Scan(source)
	statements := parser.Parse(tokens)

	for _, statement := range statements {
		statement.Accept(l.Interpreter)
	}
}

func main() {
	lox := NewLox()
	lox.Main(os.Args)
}
