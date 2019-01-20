package token

const (
	ILLEGAL = "ILLEGAL"
	UNKNOWN = "UNKNOWN"
	EOF     = "EOF"
	COMMENT = "COMMENT"

	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// Delimiters
	COMMA  = ","
	ASSIGN = "="
	LBRACE = "{"
	RBRACE = "}"
	COLON  = ":"
	RARROW = "->"

	// Keywords
	CELL      = "CELL"
	COMPONENT = "COMPONENT"
	INGRESS   = "INGRESS"

	// Types
	DOCKER = "DOCKER"
	TCP    = "TCP"
	UDP    = "UDP"
	HTTP   = "HTTP"

	// Fields
	IMAGE = "IMAGE"
	PORTS = "PORTS"
	ENV   = "ENV"

	// HTTP Methods
	GET    = "GET"
	DELETE = "DELETE"
	POST   = "POST"
	PUT    = "PUT"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func New(tokenType TokenType, char byte, line, column int) Token {
	return Token{Type: tokenType, Literal: string(char), Line: line, Column: column}
}

var keywordTable = map[string]TokenType{
	"Cell":      CELL,
	"Component": COMPONENT,
	"Ingress":   INGRESS,
	"Image":     IMAGE,
	"Docker":    DOCKER,
	"Ports":     PORTS,
	"TCP":       TCP,
	"Env":       ENV,
	"HTTP":      HTTP,
	"GET":       GET,
	"POST":      POST,
}

func LookupKeyword(keyword string) TokenType {
	if tok, ok := keywordTable[keyword]; ok {
		return tok
	}
	return UNKNOWN
}
