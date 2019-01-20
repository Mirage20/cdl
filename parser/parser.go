package parser

import (
	"fmt"
	"strconv"

	"github.com/mirage20/cdl/ast"
	"github.com/mirage20/cdl/lexer"
	"github.com/mirage20/cdl/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS       // ==
	LESSGREATER  // > or <
	SUM          // +
	PRODUCT      // *
	PREFIX       // -X or !X
	CALL         // myFunction(X)
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	nextToken    token.Token
	errors       []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.next()
	p.next()
	return p
}

func (p *Parser) ParseFile() *ast.File {
	file := &ast.File{}
	file.Cells = []*ast.Cell{}
	for p.currentToken.Type != token.EOF {
		cell := p.parseCell()
		if cell != nil {
			file.Cells = append(file.Cells, cell)
		}
		p.next()
	}
	return file
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) next() {
	p.currentToken = p.nextToken
	p.nextToken = p.l.NextToken()
	//fmt.Printf("%+v\n", p.currentToken)
}

func (p *Parser) expect(t token.TokenType) bool {
	if p.nextToken.Type == t {
		p.next()
		return true
	} else {
		p.errorExpect(t)
		return false
	}
}

func (p *Parser) errorExpect(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead at %d:%d", t, p.nextToken.Type, p.nextToken.Line, p.nextToken.Column)
	//fmt.Println(msg)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseCell() *ast.Cell {

	if p.currentToken.Type == token.CELL {
		cell := &ast.Cell{Token: p.currentToken}
		p.expect(token.IDENTIFIER)
		cell.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
		p.expect(token.LBRACE)
		for p.nextToken.Type != token.RBRACE && p.nextToken.Type != token.EOF {
			p.next()
			if p.currentToken.Type == token.COMPONENT {
				cell.Components = append(cell.Components, p.parseComponent())

			} else if p.currentToken.Type == token.INGRESS {
				cell.Ingresses = append(cell.Ingresses, p.parseIngress())
			} else {
				p.errors = append(p.errors, fmt.Sprintf("Expected one of (Component|Ingress) but got=%q", p.currentToken.Type))
			}
		}
		p.expect(token.RBRACE)
		return cell
	} else {
		return nil
	}
}

func (p *Parser) parseComponent() *ast.Component {

	if p.currentToken.Type == token.COMPONENT {
		component := &ast.Component{Token: p.currentToken}
		p.expect(token.IDENTIFIER)
		component.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
		p.expect(token.LBRACE)
		for p.nextToken.Type != token.RBRACE && p.nextToken.Type != token.EOF {
			p.next()
			switch p.currentToken.Type {
			case token.IMAGE:
				component.Image = p.parseImage()
			case token.PORTS:
				component.Ports = p.parsePorts()
			case token.ENV:
				component.Env = p.parseEnv()
			default:
				p.errors = append(p.errors, fmt.Sprintf("Unknown component field at %d:%d", p.currentToken.Line, p.currentToken.Column))
			}
		}
		p.expect(token.RBRACE)
		return component
	} else {
		return nil
	}
}
func (p *Parser) parseImage() *ast.Image {
	if p.currentToken.Type == token.IMAGE {
		image := &ast.Image{}
		p.expect(token.COLON)
		p.expect(token.DOCKER)
		image.Type = p.currentToken.Literal
		p.expect(token.ASSIGN)
		p.expect(token.STRING)
		image.Name = p.currentToken.Literal
		return image
	} else {
		return nil
	}
}

func (p *Parser) parsePorts() []*ast.Port {
	if p.currentToken.Type == token.PORTS {
		var ports []*ast.Port
		p.expect(token.LBRACE)
		for p.nextToken.Type != token.RBRACE && p.nextToken.Type != token.EOF {
			p.next()
			switch p.currentToken.Type {
			case token.TCP:
				port := &ast.Port{Type: p.currentToken.Literal}
				p.expect(token.NUMBER)
				port.HostPort = p.toInt(p.currentToken.Literal)
				p.expect(token.RARROW)
				p.expect(token.NUMBER)
				port.ContainerPort = p.toInt(p.currentToken.Literal)
				ports = append(ports, port)
			default:
				p.errors = append(p.errors, fmt.Sprintf("Unknown L4 protocol %s at %d:%d", p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column))
			}
		}
		p.expect(token.RBRACE)
		return ports
	} else {
		return nil
	}
}

func (p *Parser) parseEnv() []*ast.Env {
	if p.currentToken.Type == token.ENV {
		var envs []*ast.Env
		p.expect(token.LBRACE)
		for p.nextToken.Type != token.RBRACE && p.nextToken.Type != token.EOF {
			p.expect(token.STRING)
			env := &ast.Env{Key: p.currentToken.Literal}
			p.expect(token.COLON)
			p.expect(token.CELL)
			env.Type = p.currentToken.Literal
			p.expect(token.ASSIGN)
			p.expect(token.IDENTIFIER)
			env.Value = p.currentToken.Literal
			envs = append(envs, env)
		}
		p.expect(token.RBRACE)
		return envs
	} else {
		return nil
	}
}

func (p *Parser) parseIngress() *ast.Ingress {

	if p.currentToken.Type == token.INGRESS {
		ingress := &ast.Ingress{Token: p.currentToken}
		p.expect(token.COLON)
		p.expect(token.HTTP)
		ingress.Type = p.currentToken.Literal
		p.expect(token.LBRACE)
		for p.nextToken.Type != token.RBRACE && p.nextToken.Type != token.EOF {
			p.expect(token.STRING)
			route := &ast.Route{Context: p.currentToken.Literal}
			p.expect(token.RARROW)
			p.expect(token.IDENTIFIER)
			route.Backend = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
			route.Resources = p.parseResources()
			ingress.Routes = append(ingress.Routes, route)
		}
		p.expect(token.RBRACE)
		return ingress
	} else {
		return nil
	}
}

func (p *Parser) parseResources() []*ast.Resource {
	if p.currentToken.Type == token.IDENTIFIER {
		var resources []*ast.Resource
		p.expect(token.LBRACE)
		for p.nextToken.Type != token.RBRACE && p.nextToken.Type != token.EOF {
			p.next()
			switch p.currentToken.Type {
			case token.GET:
				resource := &ast.Resource{Method: p.currentToken.Literal}
				p.expect(token.STRING)
				resource.Path = p.currentToken.Literal
				resources = append(resources, resource)
			case token.POST:
				resource := &ast.Resource{Method: p.currentToken.Literal}
				p.expect(token.STRING)
				resource.Path = p.currentToken.Literal
				resources = append(resources, resource)
			default:
				p.errors = append(p.errors, fmt.Sprintf("Unknown HTTP verb %s at %d:%d", p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column))
			}
		}
		p.expect(token.RBRACE)
		return resources
	} else {
		return nil
	}
}

func (p *Parser) toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("Cannot convert %s to int at %d:%d", s, p.currentToken.Line, p.currentToken.Column))
		return 0
	}
	return n
}
