package parser

import (
	"golox/pkg/lox/ast"
	"golox/pkg/lox/token"
	"testing"
)

func TestParser_ParseDeclaration_ClassDeclaration(t *testing.T) {
	t.Skip("not implemented")
}

func TestParser_ParseDeclaration_FunctionDeclaration(t *testing.T) {
	t.Skip("not implemented")
}

func TestParser_ParseDeclaration_VarDeclaration(t *testing.T) {
	tokens := []token.Token{
		{token.VAR, "var", nil, 1},
		{token.IDENTIFIER, "name", nil, 1},
		{token.EQUAL, "=", nil, 1},
		{token.NUMBER, "42", 42, 1},
		{token.SEMICOLON, ";", nil, 1},
		{token.EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	declaration := parser.ParseDeclaration()

	_, ok := declaration.(ast.Var)
	if !ok {
		t.Fatal("expected a 'Var' statement")
	}

	if declaration.(ast.Var).Name.Lexeme != "name" {
		t.Fatal("expected identifier 'name'")
	}

	_, ok = declaration.(ast.Var).Initializer.(ast.Literal)
	if !ok {
		t.Fatal("expected a 'Literal' expression")
	}
}

func TestParser_ParseStatement_ExpressionStatement(t *testing.T) {
	tokens := []token.Token{
		{token.NUMBER, "42", 42, 1},
		{token.SEMICOLON, ";", nil, 1},
		{token.EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(ast.ExpressionStatement)
	if !ok {
		t.Fatal("expected an 'Expression' statement")
	}
}

func TestParser_ParseStatement_ForStatement(t *testing.T) {
	tokens := []token.Token{
		{token.FOR, "for", nil, 1},
		{token.LEFT_PAREN, "(", nil, 1},
		{token.VAR, "var", nil, 1},
		{token.IDENTIFIER, "i", nil, 1},
		{token.EQUAL, "=", nil, 1},
		{token.NUMBER, "0", 0, 1},
		{token.SEMICOLON, ";", nil, 1},
		{token.IDENTIFIER, "i", nil, 1},
		{token.LESS, "<", nil, 1},
		{token.NUMBER, "5", 5, 1},
		{token.SEMICOLON, ";", nil, 1},
		{token.IDENTIFIER, "i", nil, 1},
		{token.EQUAL, "=", nil, 1},
		{token.IDENTIFIER, "i", nil, 1},
		{token.PLUS, "+", nil, 1},
		{token.NUMBER, "1", 1, 1},
		{token.RIGHT_PAREN, ")", nil, 1},
		{token.LEFT_BRACE, "{", nil, 1},
		{token.RIGHT_BRACE, "}", nil, 1},
		{token.EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(ast.Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}

	if len(statement.(ast.Block).Statements) != 2 {
		t.Fatal("expected four statements")
	}

	_, ok = statement.(ast.Block).Statements[0].(ast.Var)
	if !ok {
		t.Fatal("expected a 'Var' statement")
	}

	_, ok = statement.(ast.Block).Statements[1].(ast.While)
	if !ok {
		t.Fatal("expected a 'While' statement")
	}
}

func TestParser_ParseStatement_IfStatement(t *testing.T) {
	tokens := []token.Token{
		{token.IF, "if", nil, 1},
		{token.LEFT_PAREN, "(", nil, 1},
		{token.TRUE, "true", true, 1},
		{token.RIGHT_PAREN, ")", nil, 1},
		{token.LEFT_BRACE, "{", nil, 1},
		{token.RIGHT_BRACE, "}", nil, 1},
		{token.ELSE, "else", nil, 1},
		{token.LEFT_BRACE, "{", nil, 1},
		{token.RIGHT_BRACE, "}", nil, 1},
		{token.EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(ast.If)
	if !ok {
		t.Fatal("expected an 'If' statement")
	}

	_, ok = statement.(ast.If).Condition.(ast.Literal)
	if !ok {
		t.Fatal("expected a 'Literal' expression")
	}

	_, ok = statement.(ast.If).ThenBranch.(ast.Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}

	_, ok = statement.(ast.If).ElseBranch.(ast.Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}
}

func TestParser_ParseStatement_PrintStatement(t *testing.T) {
	tokens := []token.Token{
		{token.PRINT, "print", nil, 1},
		{token.NUMBER, "42", 42, 1},
		{token.SEMICOLON, ";", nil, 1},
		{token.EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(ast.Print)
	if !ok {
		t.Fatal("expected a 'Print' statement")
	}

	_, ok = statement.(ast.Print).Expr.(ast.Literal)
	if !ok {
		t.Fatal("expected a 'Literal' expression")
	}
}

func TestParser_ParseStatement_ReturnStatement(t *testing.T) {
	t.Skip("not implemented")
}

func TestParser_ParseStatement_WhileStatement(t *testing.T) {
	tokens := []token.Token{
		{token.WHILE, "if", nil, 1},
		{token.LEFT_PAREN, "(", nil, 1},
		{token.TRUE, "true", true, 1},
		{token.RIGHT_PAREN, ")", nil, 1},
		{token.LEFT_BRACE, "{", nil, 1},
		{token.RIGHT_BRACE, "}", nil, 1},
		{token.EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(ast.While)
	if !ok {
		t.Fatal("expected an 'While' statement")
	}

	_, ok = statement.(ast.While).Condition.(ast.Literal)
	if !ok {
		t.Fatal("expected a 'Literal' expression")
	}

	_, ok = statement.(ast.While).Body.(ast.Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}
}

func TestParser_ParseStatement_BlockStatement(t *testing.T) {
	tokens := []token.Token{
		{token.LEFT_BRACE, "{", nil, 1},
		{token.PRINT, "print", nil, 1},
		{token.NUMBER, "1", 1, 1},
		{token.SEMICOLON, ";", nil, 1},
		{token.PRINT, "print", nil, 1},
		{token.NUMBER, "2", 2, 1},
		{token.SEMICOLON, ";", nil, 1},
		{token.RIGHT_BRACE, "}", nil, 1},
		{token.EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(ast.Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}

	if len(statement.(ast.Block).Statements) != 2 {
		t.Fatal("expected two statements")
	}

	_, ok = statement.(ast.Block).Statements[0].(ast.Print)
	if !ok {
		t.Fatal("expected a 'Print' statement")
	}

	_, ok = statement.(ast.Block).Statements[1].(ast.Print)
	if !ok {
		t.Fatal("expected a 'Print' statement")
	}
}
