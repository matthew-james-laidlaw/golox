package lox

import (
	"testing"
)

func TestParser_ParseDeclaration_ClassDeclaration(t *testing.T) {
	t.Skip("not implemented")
}

func TestParser_ParseDeclaration_FunctionDeclaration(t *testing.T) {
	t.Skip("not implemented")
}

func TestParser_ParseDeclaration_VarDeclaration(t *testing.T) {
	tokens := []Token{
		Token{VAR, "var", nil, 1},
		Token{IDENTIFIER, "name", nil, 1},
		Token{EQUAL, "=", nil, 1},
		Token{NUMBER, "42", 42, 1},
		Token{SEMICOLON, ";", nil, 1},
		Token{EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	declaration := parser.ParseDeclaration()

	_, ok := declaration.(Var)
	if !ok {
		t.Fatal("expected a 'Var' statement")
	}

	if declaration.(Var).Name.Lexeme != "name" {
		t.Fatal("expected identifier 'name'")
	}

	_, ok = declaration.(Var).Initializer.(Literal)
	if !ok {
		t.Fatal("expected a 'Literal' expression")
	}
}

func TestParser_ParseStatement_ExpressionStatement(t *testing.T) {
	tokens := []Token{
		Token{NUMBER, "42", 42, 1},
		Token{SEMICOLON, ";", nil, 1},
		Token{EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(ExpressionStatement)
	if !ok {
		t.Fatal("expected an 'Expression' statement")
	}
}

func TestParser_ParseStatement_ForStatement(t *testing.T) {
	tokens := []Token{
		Token{FOR, "for", nil, 1},
		Token{LEFT_PAREN, "(", nil, 1},
		Token{VAR, "var", nil, 1},
		Token{IDENTIFIER, "i", nil, 1},
		Token{EQUAL, "=", nil, 1},
		Token{NUMBER, "0", 0, 1},
		Token{SEMICOLON, ";", nil, 1},
		Token{IDENTIFIER, "i", nil, 1},
		Token{LESS, "<", nil, 1},
		Token{NUMBER, "5", 5, 1},
		Token{SEMICOLON, ";", nil, 1},
		Token{IDENTIFIER, "i", nil, 1},
		Token{EQUAL, "=", nil, 1},
		Token{IDENTIFIER, "i", nil, 1},
		Token{PLUS, "+", nil, 1},
		Token{NUMBER, "1", 1, 1},
		Token{RIGHT_PAREN, ")", nil, 1},
		Token{LEFT_BRACE, "{", nil, 1},
		Token{RIGHT_BRACE, "}", nil, 1},
		Token{EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}

	if len(statement.(Block).Statements) != 2 {
		t.Fatal("expected four statements")
	}

	_, ok = statement.(Block).Statements[0].(Var)
	if !ok {
		t.Fatal("expected a 'Var' statement")
	}

	_, ok = statement.(Block).Statements[1].(While)
	if !ok {
		t.Fatal("expected a 'While' statement")
	}
}

func TestParser_ParseStatement_IfStatement(t *testing.T) {
	tokens := []Token{
		Token{IF, "if", nil, 1},
		Token{LEFT_PAREN, "(", nil, 1},
		Token{TRUE, "true", true, 1},
		Token{RIGHT_PAREN, ")", nil, 1},
		Token{LEFT_BRACE, "{", nil, 1},
		Token{RIGHT_BRACE, "}", nil, 1},
		Token{ELSE, "else", nil, 1},
		Token{LEFT_BRACE, "{", nil, 1},
		Token{RIGHT_BRACE, "}", nil, 1},
		Token{EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(If)
	if !ok {
		t.Fatal("expected an 'If' statement")
	}

	_, ok = statement.(If).Condition.(Literal)
	if !ok {
		t.Fatal("expected a 'Literal' expression")
	}

	_, ok = statement.(If).ThenBranch.(Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}

	_, ok = statement.(If).ElseBranch.(Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}
}

func TestParser_ParseStatement_PrintStatement(t *testing.T) {
	tokens := []Token{
		Token{PRINT, "print", nil, 1},
		Token{NUMBER, "42", 42, 1},
		Token{SEMICOLON, ";", nil, 1},
		Token{EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(Print)
	if !ok {
		t.Fatal("expected a 'Print' statement")
	}

	_, ok = statement.(Print).Expr.(Literal)
	if !ok {
		t.Fatal("expected a 'Literal' expression")
	}
}

func TestParser_ParseStatement_ReturnStatement(t *testing.T) {
	t.Skip("not implemented")
}

func TestParser_ParseStatement_WhileStatement(t *testing.T) {
	tokens := []Token{
		Token{WHILE, "if", nil, 1},
		Token{LEFT_PAREN, "(", nil, 1},
		Token{TRUE, "true", true, 1},
		Token{RIGHT_PAREN, ")", nil, 1},
		Token{LEFT_BRACE, "{", nil, 1},
		Token{RIGHT_BRACE, "}", nil, 1},
		Token{EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(While)
	if !ok {
		t.Fatal("expected an 'While' statement")
	}

	_, ok = statement.(While).Condition.(Literal)
	if !ok {
		t.Fatal("expected a 'Literal' expression")
	}

	_, ok = statement.(While).Body.(Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}
}

func TestParser_ParseStatement_BlockStatement(t *testing.T) {
	tokens := []Token{
		Token{LEFT_BRACE, "{", nil, 1},
		Token{PRINT, "print", nil, 1},
		Token{NUMBER, "1", 1, 1},
		Token{SEMICOLON, ";", nil, 1},
		Token{PRINT, "print", nil, 1},
		Token{NUMBER, "2", 2, 1},
		Token{SEMICOLON, ";", nil, 1},
		Token{RIGHT_BRACE, "}", nil, 1},
		Token{EOF, "", nil, 2},
	}

	parser := NewParser(tokens)
	statement := parser.ParseStatement()

	_, ok := statement.(Block)
	if !ok {
		t.Fatal("expected a 'Block' statement")
	}

	if len(statement.(Block).Statements) != 2 {
		t.Fatal("expected two statements")
	}

	_, ok = statement.(Block).Statements[0].(Print)
	if !ok {
		t.Fatal("expected a 'Print' statement")
	}

	_, ok = statement.(Block).Statements[1].(Print)
	if !ok {
		t.Fatal("expected a 'Print' statement")
	}
}
