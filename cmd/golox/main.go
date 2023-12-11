package main

import (
	"bufio"
	"fmt"
	"golox/pkg/lox"
	"golox/pkg/lox/parser"
	"log"
	"os"
)

type Lox struct {
	Interpreter *lox.Interpreter
}

func NewLox() *Lox {
	return &Lox{
		Interpreter: lox.NewInterpreter(),
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
	scanner := lox.NewScanner(source)
	tokens := scanner.ScanTokens()

	tokenParser := parser.NewParser(tokens)
	statements := tokenParser.Parse()

	for _, statement := range statements {
		statement.Accept(l.Interpreter)
	}
}

func main() {
	lox := NewLox()
	lox.Main(os.Args)
}
