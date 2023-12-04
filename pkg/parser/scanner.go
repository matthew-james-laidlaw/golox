package parser

import (
	"log"
	"strconv"
)

type Scanner struct {
	Source  string
	Tokens  []Token
	Start   int
	Current int
	Line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		Source:  source,
		Tokens:  make([]Token, 0),
		Start:   0,
		Current: 0,
		Line:    1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.IsAtEnd() {
		s.Start = s.Current
		s.ScanToken()
	}

	eof := Token{
		Type:    EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    s.Line,
	}

	s.Tokens = append(s.Tokens, eof)

	return s.Tokens
}

func (s *Scanner) IsAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) ScanToken() {
	c := s.Advance()
	switch c {
	case '(':
		s.AddToken(LEFT_PAREN)
	case ')':
		s.AddToken(RIGHT_PAREN)
	case '{':
		s.AddToken(LEFT_BRACE)
	case '}':
		s.AddToken(RIGHT_BRACE)
	case ',':
		s.AddToken(COMMA)
	case '.':
		s.AddToken(DOT)
	case '-':
		s.AddToken(MINUS)
	case '+':
		s.AddToken(PLUS)
	case ';':
		s.AddToken(SEMICOLON)
	case '*':
		s.AddToken(STAR)
	case '!':
		if s.Match('=') {
			s.AddToken(BANG_EQUAL)
		} else {
			s.AddToken(BANG)
		}
	case '=':
		if s.Match('=') {
			s.AddToken(EQUAL_EQUAL)
		} else {
			s.AddToken(BANG)
		}
	case '<':
		if s.Match('=') {
			s.AddToken(LESS_EQUAL)
		} else {
			s.AddToken(BANG)
		}
	case '>':
		if s.Match('=') {
			s.AddToken(GREATER_EQUAL)
		} else {
			s.AddToken(BANG)
		}
	case '/':
		if s.Match('/') {
			for s.Peek() != '\n' && !s.IsAtEnd() {
				s.Advance()
			}
		} else {
			s.AddToken(SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.Line++
	case '"':
		s.ScanString()
	default:
		if IsDigit(c) {
			s.ScanNumber()
		} else if IsAlpha(c) {
			s.ScanIdentifier()
		} else {
			log.Fatalln("unexpected character")
		}
	}
}

func (s *Scanner) Advance() byte {
	next := s.Source[s.Current]
	s.Current++
	return next
}

func (s *Scanner) AddToken(tokenType TokenType) {
	s.AddTokenWithValue(tokenType, nil)
}

func (s *Scanner) AddTokenWithValue(tokenType TokenType, literal interface{}) {
	text := s.Source[s.Start:s.Current]
	s.Tokens = append(s.Tokens, Token{tokenType, text, literal, s.Line})
}

func (s *Scanner) Match(expected byte) bool {
	if s.IsAtEnd() {
		return false
	}

	if s.Source[s.Current] != expected {
		return false
	}

	s.Current++

	return true
}

func (s *Scanner) Peek() byte {
	if s.IsAtEnd() {
		return '\000'
	}
	return s.Source[s.Current]
}

func (s *Scanner) ScanString() {
	for s.Peek() != '"' && !s.IsAtEnd() {
		if s.Peek() == '\n' {
			s.Line++
		}
		s.Advance()
	}

	if s.IsAtEnd() {
		log.Fatalln("unterminated string")
	}

	s.Advance()

	value := s.Source[s.Start+1 : s.Current-1]
	s.AddTokenWithValue(STRING, value)
}

func (s *Scanner) ScanNumber() {
	for IsDigit(s.Peek()) {
		s.Advance()
	}

	if s.Peek() == '.' && IsDigit(s.PeekNext()) {
		s.Advance()
		for IsDigit(s.Peek()) {
			s.Advance()
		}
	}

	value, err := strconv.ParseFloat(s.Source[s.Start:s.Current], 64)
	if err != nil {
		log.Fatalln(err)
	}

	s.AddTokenWithValue(NUMBER, value)
}

func (s *Scanner) ScanIdentifier() {
	for IsAlphaNumeric(s.Peek()) {
		s.Advance()
	}

	text := s.Source[s.Start:s.Current]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = IDENTIFIER
	}
	s.AddToken(tokenType)
}

func (s *Scanner) PeekNext() byte {
	if s.Current+1 >= len(s.Source) {
		return '\000'
	}
	return s.Source[s.Current+1]
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func IsAlphaNumeric(c byte) bool {
	return IsDigit(c) || IsAlpha(c)
}
