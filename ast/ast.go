package ast

import (
	"github.com/mirage20/cdl/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type File struct {
	Cells []*Cell
}

type Cell struct {
	Token      token.Token
	Name       *Identifier
	Components []*Component
	Ingresses  []*Ingress
}

type Component struct {
	Token token.Token
	Name  *Identifier
	Image *Image
	Ports []*Port
	Env   []*Env
}

type Image struct {
	Type string
	Name string
}

type Port struct {
	Type          string
	HostPort      int
	ContainerPort int
}

type Env struct {
	Type  string
	Key   string
	Value string
}

type Ingress struct {
	Token  token.Token
	Type   string
	Routes []*Route
}

type Route struct {
	Context   string
	Backend   *Identifier
	Resources []*Resource
}

type Resource struct {
	Method string
	Path   string
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
