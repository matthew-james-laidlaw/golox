package parser

import (
	"golox/pkg/lox/ast"
	"golox/pkg/lox/token"
)

type Parser struct {
	Tokens  []token.Token
	Current int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		Tokens:  tokens,
		Current: 0,
	}
}

func (p *Parser) Parse() []ast.Statement {
	statements := make([]ast.Statement, 0)
	for !p.IsAtEnd() {
		statements = append(statements, p.ParseDeclaration())
	}
	return statements
}

func (p *Parser) ParseDeclaration() ast.Statement {
	if p.Match(token.FUN) {
		return p.ParseFunctionDeclaration()
	}
	if p.Match(token.VAR) {
		return p.ParseVarDeclaration()
	}
	return p.ParseStatement()
}

func (p *Parser) ParseFunctionDeclaration() ast.Statement {
	name := p.Consume(token.IDENTIFIER, "expect function name")
	p.Consume(token.LEFT_PAREN, "expect '(' after function name")
	parameters := make([]token.Token, 0)
	if !p.Check(token.RIGHT_PAREN) {
		parameters = append(parameters, p.Consume(token.IDENTIFIER, "expect parameter name"))
		for p.Match(token.COMMA) {
			parameters = append(parameters, p.Consume(token.IDENTIFIER, "expect parameter name"))
		}
	}
	p.Consume(token.RIGHT_PAREN, "expect ')' after parameters")

	p.Consume(token.LEFT_BRACE, "expect '{' before function body")
	body := p.ParseBlock()

	return ast.Function{
		Name:   name,
		Params: parameters,
		Body:   body,
	}
}

func (p *Parser) ParseVarDeclaration() ast.Statement {
	name := p.Consume(token.IDENTIFIER, "expect variable name")

	var initializer ast.Expression
	if p.Match(token.EQUAL) {
		initializer = p.ParseExpression()
	}

	p.Consume(token.SEMICOLON, "expect ';' after variable declaration")

	return ast.Var{
		Name:        name,
		Initializer: initializer,
	}
}

func (p *Parser) ParseStatement() ast.Statement {
	if p.Match(token.IF) {
		return p.ParseIfStatement()
	}
	if p.Match(token.PRINT) {
		return p.ParsePrintStatement()
	}
	if p.Match(token.RETURN) {
		return p.ParseReturn()
	}
	if p.Match(token.WHILE) {
		return p.ParseWhileStatement()
	}
	if p.Match(token.LEFT_BRACE) {
		return ast.Block{
			Statements: p.ParseBlock(),
		}
	}
	if p.Match(token.FOR) {
		return p.ParseForStatement()
	}

	return p.ParseExpressionStatement()
}

func (p *Parser) ParseReturn() ast.Statement {
	_ = p.Previous()
	var value ast.Expression
	if !p.Check(token.SEMICOLON) {
		value = p.ParseExpression()
	}
	p.Consume(token.SEMICOLON, "expected ';' after return value")

	return ast.Return{
		Value: value,
	}
}

func (p *Parser) ParseWhileStatement() ast.Statement {
	p.Consume(token.LEFT_PAREN, "expect '(' after 'while'")
	condition := p.ParseExpression()
	p.Consume(token.RIGHT_PAREN, "expect ')' after while condition")
	body := p.ParseStatement()

	return ast.While{
		Condition: condition,
		Body:      body,
	}
}

func (p *Parser) ParseIfStatement() ast.Statement {
	p.Consume(token.LEFT_PAREN, "expect '(' after 'if'")
	condition := p.ParseExpression()
	p.Consume(token.RIGHT_PAREN, "expect ')' after if condition")

	thenBranch := p.ParseStatement()
	var elseBranch ast.Statement

	if p.Match(token.ELSE) {
		elseBranch = p.ParseStatement()
	}

	return ast.If{
		Condition:  condition,
		ThenBranch: thenBranch,
		ElseBranch: elseBranch,
	}
}

func (p *Parser) ParseForStatement() ast.Statement {
	p.Consume(token.LEFT_PAREN, "expect '(' after 'for'")

	var initializer ast.Statement
	if p.Match(token.SEMICOLON) {
		initializer = nil
	} else if p.Match(token.VAR) {
		initializer = p.ParseVarDeclaration()
	} else {
		initializer = p.ParseExpressionStatement()
	}

	var condition ast.Expression
	if !p.Check(token.SEMICOLON) {
		condition = p.ParseExpression()
	}
	p.Consume(token.SEMICOLON, "expect ';' after loop condition")

	var increment ast.Expression
	if !p.Check(token.RIGHT_PAREN) {
		increment = p.ParseExpression()
	}
	p.Consume(token.RIGHT_PAREN, "expect ')' after for clauses")

	body := p.ParseStatement()

	if increment != nil {
		body = ast.Block{
			Statements: []ast.Statement{body, increment},
		}
	}

	if condition == nil {
		condition = ast.Literal{
			Value: true,
		}
	}

	body = ast.While{
		Condition: condition,
		Body:      body,
	}

	if initializer != nil {
		body = ast.Block{
			Statements: []ast.Statement{initializer, body},
		}
	}

	return body
}

func (p *Parser) ParseBlock() []ast.Statement {
	statements := make([]ast.Statement, 0)
	for !p.Check(token.RIGHT_BRACE) && !p.IsAtEnd() {
		statements = append(statements, p.ParseDeclaration())
	}
	p.Consume(token.RIGHT_BRACE, "expect '}' after block")
	return statements
}

func (p *Parser) ParsePrintStatement() ast.Statement {
	value := p.ParseExpression()
	p.Consume(token.SEMICOLON, "expect ';' after value")

	return ast.Print{
		Expr: value,
	}
}

func (p *Parser) ParseExpressionStatement() ast.Statement {
	expr := p.ParseExpression()
	p.Consume(token.SEMICOLON, "expect ';' after expression")

	return ast.ExpressionStatement{
		Expr: expr,
	}
}

func (p *Parser) ParseExpression() ast.Expression {
	return p.ParseAssignment()
}

func (p *Parser) ParseAssignment() ast.Expression {
	expr := p.ParseOr()

	if p.Match(token.EQUAL) {
		value := p.ParseAssignment()

		if _, ok := expr.(ast.Variable); ok {
			name := expr.(ast.Variable).Name
			return ast.Assignment{
				Name:  name,
				Value: value,
			}
		}

		panic("invalid assignment target")
	}

	return expr
}

func (p *Parser) ParseOr() ast.Statement {
	expr := p.ParseAnd()

	for p.Match(token.OR) {
		operator := p.Previous()
		right := p.ParseAnd()
		expr = ast.Logical{
			Operation: operator,
			Left:      expr,
			Right:     right,
		}
	}

	return expr
}

func (p *Parser) ParseAnd() ast.Statement {
	expr := p.ParseEquality()

	for p.Match(token.AND) {
		operator := p.Previous()
		right := p.ParseEquality()
		expr = ast.Logical{
			Operation: operator,
			Left:      expr,
			Right:     right,
		}
	}

	return expr
}

func (p *Parser) ParseEquality() ast.Expression {
	expr := p.ParseComparison()

	for p.Match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.Previous()
		right := p.ParseComparison()
		expr = ast.Binary{
			Operation: operator,
			Left:      expr,
			Right:     right,
		}
	}

	return expr
}

func (p *Parser) Match(types ...token.TokenType) bool {
	for _, typ := range types {
		if p.Check(typ) {
			p.Advance()
			return true
		}
	}
	return false
}

func (p *Parser) Check(typ token.TokenType) bool {
	if p.IsAtEnd() {
		return false
	}
	return p.Peek().Type == typ
}

func (p *Parser) Advance() token.Token {
	if !p.IsAtEnd() {
		p.Current++
	}
	return p.Previous()
}

func (p *Parser) IsAtEnd() bool {
	return p.Peek().Type == token.EOF
}

func (p *Parser) Peek() token.Token {
	return p.Tokens[p.Current]
}

func (p *Parser) Previous() token.Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) ParseComparison() ast.Expression {
	expr := p.ParseTerm()

	for p.Match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.Previous()
		right := p.ParseTerm()
		expr = ast.Binary{
			Operation: operator,
			Left:      expr,
			Right:     right,
		}
	}

	return expr
}

func (p *Parser) ParseTerm() ast.Expression {
	expr := p.ParseFactor()

	for p.Match(token.MINUS, token.PLUS) {
		operator := p.Previous()
		right := p.ParseFactor()
		expr = ast.Binary{
			Operation: operator,
			Left:      expr,
			Right:     right,
		}
	}

	return expr
}

func (p *Parser) ParseFactor() ast.Expression {
	expr := p.ParseUnary()

	for p.Match(token.SLASH, token.STAR) {
		operator := p.Previous()
		right := p.ParseUnary()
		expr = ast.Binary{
			Operation: operator,
			Left:      expr,
			Right:     right,
		}
	}

	return expr
}

func (p *Parser) ParseUnary() ast.Expression {
	if p.Match(token.BANG, token.MINUS) {
		operator := p.Previous()
		right := p.ParseUnary()
		return ast.Unary{
			Operation: operator,
			Operand:   right,
		}
	}

	return p.ParseCall()
}

func (p *Parser) ParseCall() ast.Expression {
	expression := p.ParsePrimary()

	for {
		if p.Match(token.LEFT_PAREN) {
			expression = p.FinishCall(expression)
		} else {
			break
		}
	}

	return expression
}

func (p *Parser) FinishCall(callee ast.Expression) ast.Expression {
	arguments := make([]ast.Expression, 0)
	if !p.Check(token.RIGHT_PAREN) {
		arguments = append(arguments, p.ParseExpression())
		for p.Match(token.COMMA) {
			arguments = append(arguments, p.ParseExpression())
		}
	}

	_ = p.Consume(token.RIGHT_PAREN, "expect ')' after arguments")

	return ast.Call{
		Callee:    callee,
		Arguments: arguments,
	}
}

func (p *Parser) ParsePrimary() ast.Expression {
	if p.Match(token.FALSE) {
		return ast.Literal{
			Value: false,
		}
	}
	if p.Match(token.TRUE) {
		return ast.Literal{
			Value: true,
		}
	}
	if p.Match(token.NIL) {
		return ast.Literal{
			Value: nil,
		}
	}

	if p.Match(token.NUMBER, token.STRING) {
		return ast.Literal{Value: p.Previous().Literal}
	}

	if p.Match(token.IDENTIFIER) {
		return ast.Variable{
			Name: p.Previous(),
		}
	}

	if p.Match(token.LEFT_PAREN) {
		expr := p.ParseExpression()
		p.Consume(token.RIGHT_PAREN, "expect ')' after expression")
		return ast.Grouping{
			Expr: expr,
		}
	}

	panic("not supposed to be here!")
}

func (p *Parser) Consume(typ token.TokenType, message string) token.Token {
	if p.Check(typ) {
		return p.Advance()
	}

	panic(message)
}
