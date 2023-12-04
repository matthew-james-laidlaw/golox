package main

import (
	"bufio"
	"fmt"
	"golox/pkg/parser"
	"log"
	"os"
)

type Lox struct {
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
	scanner := parser.NewScanner(source)
	tokens := scanner.ScanTokens()
	for token, _ := range tokens {
		fmt.Printf("%+v\n", token)
	}
}

func main() {
	lox := Lox{}
	lox.Main(os.Args)
}
