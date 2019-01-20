package lexer

import (
	"github.com/mirage20/cdl/token"
)

type Lexer struct {
	input   string
	current int  // next index of the reading position
	next    int  // next index of the reading position
	char    byte // current char
	line    int
	column  int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
	return l
}

func (l *Lexer) readChar() {
	l.char = l.peekChar()
	l.current = l.next
	l.next += 1
	if l.char == '\n' {
		l.line++
		l.column = 0
	} else if l.char == '\t' {
		l.column += 4 // Assume tab size is 4
	} else {
		l.column++
	}
	//fmt.Printf("%q %d:%d\n", l.char, l.line, l.column)
}

func (l *Lexer) peekChar() byte {
	if l.next >= len(l.input) {
		return 0
	} else {
		return l.input[l.next]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespaceAndComments()
	l.readChar()
	switch l.char {
	case '=':
		// 	if l.peekChar() == '=' {
		// 		ch := l.ch
		// 		l.readChar()
		// 		tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		// 	} else {
		tok = token.New(token.ASSIGN, l.char, l.line, l.column)
		// 	}
		// case '(':
		// 	tok = newToken(token.LPAREN, l.ch)
		// case ')':
		// 	tok = newToken(token.RPAREN, l.ch)
		// case '+':
		// 	tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.char, l.line, l.column)
	case '}':
		tok = token.New(token.RBRACE, l.char, l.line, l.column)
		// case '-':
		// 	tok = newToken(token.MINUS, l.ch)
	case '-':
		if l.peekChar() == '>' {
			ch := l.char
			l.readChar()
			tok.Literal = string(ch) + string(l.char)
			tok.Type = token.RARROW
			tok.Line = l.line
			tok.Column = l.column
		} else {

		}

		// case '/':
		// 	tok = newToken(token.SLASH, l.ch)
		// case '*':
		// 	tok = newToken(token.ASTERISK, l.ch)
		// case '<':
		// 	tok = newToken(token.LT, l.ch)
		// case '>':
		// 	tok = newToken(token.GT, l.ch)
	case ':':
		tok = token.New(token.COLON, l.char, l.line, l.column)
	case ',':
		tok = token.New(token.COMMA, l.char, l.line, l.column)
	case '"':
		tok.Literal = l.readString()
		tok.Type = token.STRING
		tok.Line = l.line
		tok.Column = l.column
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.line
		tok.Column = l.column
	default:
		if isLowerCaseLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.IDENTIFIER
			tok.Line = l.line
			tok.Column = l.column
		} else if isUpperCaseLetter(l.char) {
			tok.Literal = l.readKeyword()
			tok.Type = token.LookupKeyword(tok.Literal)
			tok.Line = l.line
			tok.Column = l.column
		} else if isNumber(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.NUMBER
			tok.Line = l.line
			tok.Column = l.column
		} else {
			tok = token.New(token.ILLEGAL, l.char, l.line, l.column)
		}
	}
	return tok
}

func (l *Lexer) skipWhitespaceAndComments() {
	for {
		nextChar := l.peekChar()
		if isWhitespace(nextChar) {
			l.skipWhitespace()
			continue
		} else if nextChar == '#' {
			l.skipLineComment()
			continue
		}
		break
	}
}

func (l *Lexer) skipWhitespace() {
	nextChar := l.peekChar()
	for isWhitespace(nextChar) {
		l.readChar()
		nextChar = l.peekChar()
	}
}

func (l *Lexer) skipLineComment() {
	nextChar := l.peekChar()
	if nextChar == '#' {
		for nextChar != '\n' {
			l.readChar()
			nextChar = l.peekChar()
		}
	}
}

func (l *Lexer) readIdentifier() string {
	current := l.current
	for isIdentifierChar(l.peekChar()) {
		l.readChar()
	}
	return l.input[current:l.next]
}

func (l *Lexer) readKeyword() string {
	current := l.current
	for isKeywordChar(l.peekChar()) {
		l.readChar()
	}
	return l.input[current:l.next]
}

func (l *Lexer) readString() string {
	l.readChar()
	start := l.current
	for isAlphanumeric(l.char) {
		l.readChar()
	}
	end := l.current
	return l.input[start:end]
}

func (l *Lexer) readNumber() string {
	current := l.current
	for isNumber(l.peekChar()) {
		l.readChar()
	}
	return l.input[current:l.next]
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isIdentifierChar(ch byte) bool {
	return isLowerCaseLetter(ch) || ch == '-' || isNumber(ch)
}

func isKeywordChar(ch byte) bool {
	return isUpperCaseLetter(ch) || isLowerCaseLetter(ch)
}

func isAlphanumeric(ch byte) bool {
	return isLowerCaseLetter(ch) || isUpperCaseLetter(ch) || isNumber(ch) || ch == '_' || ch == '/' || ch == '.' || ch == '-'
}

func isLowerCaseLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z'
}

func isUpperCaseLetter(ch byte) bool {
	return 'A' <= ch && ch <= 'Z'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// func newToken(tokenType token.TokenType, ch byte) token.Token {
// 	return token.Token{Type: tokenType, Literal: string(ch)}
// }

// func (l *Lexer) peekChar() byte {
// 	if l.readPosition >= len(l.input) {
// 		return 0
// 	} else {
// 		return l.input[l.readPosition]
// 	}
// }
