package lox

type Parser struct {
	Tokens  []Token
	Current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		Tokens:  tokens,
		Current: 0,
	}
}

func (p *Parser) Parse() []Statement {
	statements := make([]Statement, 0)
	for !p.IsAtEnd() {
		statements = append(statements, p.ParseDeclaration())
	}
	return statements
}

func (p *Parser) ParseDeclaration() Statement {
	if p.Match(VAR) {
		return p.ParseVarDeclaration()
	}
	return p.ParseStatement()
}

func (p *Parser) ParseVarDeclaration() Statement {
	name := p.Consume(IDENTIFIER, "expect variable name")

	var initializer Expression
	if p.Match(EQUAL) {
		initializer = p.ParseExpression()
	}

	p.Consume(SEMICOLON, "expect ';' after variable declaration")
	return Var{name, initializer}
}

func (p *Parser) ParseStatement() Statement {
	if p.Match(IF) {
		return p.ParseIfStatement()
	}
	if p.Match(PRINT) {
		return p.ParsePrintStatement()
	}
	if p.Match(WHILE) {
		return p.ParseWhileStatement()
	}
	if p.Match(LEFT_BRACE) {
		return Block{p.ParseBlock()}
	}
	if p.Match(FOR) {
		return p.ParseForStatement()
	}

	return p.ParseExpressionStatement()
}

func (p *Parser) ParseWhileStatement() Statement {
	p.Consume(LEFT_PAREN, "expect '(' after 'while'")
	condition := p.ParseExpression()
	p.Consume(RIGHT_PAREN, "expect ')' after while condition")
	body := p.ParseStatement()

	return While{condition, body}
}

func (p *Parser) ParseIfStatement() Statement {
	p.Consume(LEFT_PAREN, "expect '(' after 'if'")
	condition := p.ParseExpression()
	p.Consume(RIGHT_PAREN, "expect ')' after if condition")

	thenBranch := p.ParseStatement()
	var elseBranch Statement

	if p.Match(ELSE) {
		elseBranch = p.ParseStatement()
	}

	return If{condition, thenBranch, elseBranch}
}

func (p *Parser) ParseForStatement() Statement {
	p.Consume(LEFT_PAREN, "expect '(' after 'for'")

	var initializer Statement
	if p.Match(SEMICOLON) {
		initializer = nil
	} else if p.Match(VAR) {
		initializer = p.ParseVarDeclaration()
	} else {
		initializer = p.ParseExpressionStatement()
	}

	var condition Expression
	if !p.Check(SEMICOLON) {
		condition = p.ParseExpression()
	}
	p.Consume(SEMICOLON, "expect ';' after loop condition")

	var increment Expression
	if !p.Check(RIGHT_PAREN) {
		increment = p.ParseExpression()
	}
	p.Consume(RIGHT_PAREN, "expect ')' after for clauses")

	body := p.ParseStatement()

	if increment != nil {
		body = Block{[]Statement{body, increment}}
	}

	if condition == nil {
		condition = Literal{true}
	}

	body = While{condition, body}

	if initializer != nil {
		body = Block{[]Statement{initializer, body}}
	}

	return body
}

func (p *Parser) ParseBlock() []Statement {
	statements := make([]Statement, 0)
	for !p.Check(RIGHT_BRACE) && !p.IsAtEnd() {
		statements = append(statements, p.ParseDeclaration())
	}
	p.Consume(RIGHT_BRACE, "expect '}' after block")
	return statements
}

func (p *Parser) ParsePrintStatement() Statement {
	value := p.ParseExpression()
	p.Consume(SEMICOLON, "expect ';' after value")
	return Print{value}
}

func (p *Parser) ParseExpressionStatement() Statement {
	expr := p.ParseExpression()
	p.Consume(SEMICOLON, "expect ';' after expression")
	return ExpressionStatement{expr}
}

func (p *Parser) ParseExpression() Expression {
	return p.ParseAssignment()
}

func (p *Parser) ParseAssignment() Expression {
	expr := p.ParseOr()

	if p.Match(EQUAL) {
		value := p.ParseAssignment()

		if _, ok := expr.(Variable); ok {
			name := expr.(Variable).Name
			return Assignment{name, value}
		}

		panic("invalid assignment target")
	}

	return expr
}

func (p *Parser) ParseOr() Statement {
	expr := p.ParseAnd()

	for p.Match(OR) {
		operator := p.Previous()
		right := p.ParseAnd()
		expr = Logical{operator, expr, right}
	}

	return expr
}

func (p *Parser) ParseAnd() Statement {
	expr := p.ParseEquality()

	for p.Match(AND) {
		operator := p.Previous()
		right := p.ParseEquality()
		expr = Logical{operator, expr, right}
	}

	return expr
}

func (p *Parser) ParseEquality() Expression {
	expr := p.ParseComparison()

	for p.Match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.Previous()
		right := p.ParseComparison()
		expr = Binary{operator, expr, right}
	}

	return expr
}

func (p *Parser) Match(types ...TokenType) bool {
	for _, typ := range types {
		if p.Check(typ) {
			p.Advance()
			return true
		}
	}
	return false
}

func (p *Parser) Check(typ TokenType) bool {
	if p.IsAtEnd() {
		return false
	}
	return p.Peek().Type == typ
}

func (p *Parser) Advance() Token {
	if !p.IsAtEnd() {
		p.Current++
	}
	return p.Previous()
}

func (p *Parser) IsAtEnd() bool {
	return p.Peek().Type == EOF
}

func (p *Parser) Peek() Token {
	return p.Tokens[p.Current]
}

func (p *Parser) Previous() Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) ParseComparison() Expression {
	expr := p.ParseTerm()

	for p.Match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.Previous()
		right := p.ParseTerm()
		expr = Binary{operator, expr, right}
	}

	return expr
}

func (p *Parser) ParseTerm() Expression {
	expr := p.ParseFactor()

	for p.Match(MINUS, PLUS) {
		operator := p.Previous()
		right := p.ParseFactor()
		expr = Binary{operator, expr, right}
	}

	return expr
}

func (p *Parser) ParseFactor() Expression {
	expr := p.ParseUnary()

	for p.Match(SLASH, STAR) {
		operator := p.Previous()
		right := p.ParseUnary()
		expr = Binary{operator, expr, right}
	}

	return expr
}

func (p *Parser) ParseUnary() Expression {
	if p.Match(BANG, MINUS) {
		operator := p.Previous()
		right := p.ParseUnary()
		return Unary{operator, right}
	}
	return p.ParsePrimary()
}

func (p *Parser) ParsePrimary() Expression {
	if p.Match(FALSE) {
		return Literal{false}
	}
	if p.Match(TRUE) {
		return Literal{true}
	}
	if p.Match(NIL) {
		return Literal{nil}
	}

	if p.Match(NUMBER, STRING) {
		return Literal{p.Previous().Literal}
	}

	if p.Match(IDENTIFIER) {
		return Variable{p.Previous()}
	}

	if p.Match(LEFT_PAREN) {
		expr := p.ParseExpression()
		p.Consume(RIGHT_PAREN, "expect ')' after expression")
		return Grouping{expr}
	}

	panic("not supposed to be here!")
}

func (p *Parser) Consume(typ TokenType, message string) Token {
	if p.Check(typ) {
		return p.Advance()
	}

	panic(message)
}
