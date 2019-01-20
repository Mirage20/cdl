package parser

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/mirage20/cdl/ast"
	"github.com/mirage20/cdl/lexer"
	"github.com/mirage20/cdl/token"
	"testing"
)

func TestCellBlock(t *testing.T) {
	input := `



	# This is a comment

	Cell employee {
		Component employee {
			Image:Docker = "docker.io/wso2vick/sampleapp-employee"
			Ports {
				TCP 80->8080
				TCP 443->8081
			}
		}
	
		Component salary {
			Image:Docker = "docker.io/wso2vick/sampleapp-salary"
			Ports {
				 TCP 80->8080
			}
		}
	
		Ingress:HTTP {
			"/employee" -> employee {
				GET "/"
				POST "/"
			}
		}
	}
	
	
	
	Cell stock-options {
		Component stock {
			Image:Docker = "docker.io/wso2vick/sampleapp-stock"
			Ports {
				 TCP 80->8080
			}
		}
	
		Ingress:HTTP {
			"/stock" -> stock {
				GET "/"
			}
		}
	}
	
	
	
	Cell hr {
		Component hr {
			Image:Docker = "docker.io/wso2vick/sampleapp-hr"
			Ports {
				 TCP 80->8080
			}
			Env {
				"employeegw_url":Cell = employee
				"stockgw_url":Cell = stock-options
			}
		}
	
		Ingress:HTTP {
			"/info" -> hr {
				GET "/"
			}
		}
	}
	
	
	

`
	l := lexer.New(input)
	p := New(l)
	file := p.ParseFile()
	checkParserErrors(t, p)
	if file == nil {
		t.Fatalf("ParseFile() returned nil")
	}
	if len(file.Cells) != 3 {
		t.Fatalf("file.Cells does not contain 3 cells. got=%d",
			len(file.Cells))
	}

	tests := []struct {
		want *ast.Cell
	}{
		{
			want: &ast.Cell{
				Token: token.Token{
					Type:    token.CELL,
					Literal: "Cell",
				},
				Name: &ast.Identifier{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "employee",
					},
					Value: "employee",
				},
				Components: []*ast.Component{
					{
						Token: token.Token{
							Type:    token.COMPONENT,
							Literal: "Component",
						},
						Name: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENTIFIER,
								Literal: "employee",
							},
							Value: "employee",
						},
						Image: &ast.Image{
							Type: "Docker",
							Name: "docker.io/wso2vick/sampleapp-employee",
						},
						Ports: []*ast.Port{
							{
								Type:          "TCP",
								ContainerPort: 8080,
								HostPort:      80,
							},
							{
								Type:          "TCP",
								ContainerPort: 8081,
								HostPort:      443,
							},
						},
					},
					{
						Token: token.Token{
							Type:    token.COMPONENT,
							Literal: "Component",
						},
						Name: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENTIFIER,
								Literal: "salary",
							},
							Value: "salary",
						},
						Image: &ast.Image{
							Type: "Docker",
							Name: "docker.io/wso2vick/sampleapp-salary",
						},
						Ports: []*ast.Port{
							{
								Type:          "TCP",
								ContainerPort: 8080,
								HostPort:      80,
							},
						},
					},
				},
				Ingresses: []*ast.Ingress{
					{
						Token: token.Token{
							Type:    token.INGRESS,
							Literal: "Ingress",
						},
						Type: "HTTP",
						Routes: []*ast.Route{
							{
								Context: "/employee",
								Backend: &ast.Identifier{
									Token: token.Token{
										Type:    token.IDENTIFIER,
										Literal: "employee",
									},
									Value: "employee",
								},
								Resources: []*ast.Resource{
									{
										Path:   "/",
										Method: "GET",
									},
									{
										Path:   "/",
										Method: "POST",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			want: &ast.Cell{
				Token: token.Token{
					Type:    token.CELL,
					Literal: "Cell",
				},
				Name: &ast.Identifier{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "stock-options",
					},
					Value: "stock-options",
				},
				Components: []*ast.Component{
					{
						Token: token.Token{
							Type:    token.COMPONENT,
							Literal: "Component",
						},
						Name: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENTIFIER,
								Literal: "stock",
							},
							Value: "stock",
						},
						Image: &ast.Image{
							Type: "Docker",
							Name: "docker.io/wso2vick/sampleapp-stock",
						},
						Ports: []*ast.Port{
							{
								Type:          "TCP",
								ContainerPort: 8080,
								HostPort:      80,
							},
						},
					},
				},
				Ingresses: []*ast.Ingress{
					{
						Token: token.Token{
							Type:    token.INGRESS,
							Literal: "Ingress",
						},
						Type: "HTTP",
						Routes: []*ast.Route{
							{
								Context: "/stock",
								Backend: &ast.Identifier{
									Token: token.Token{
										Type:    token.IDENTIFIER,
										Literal: "stock",
									},
									Value: "stock",
								},
								Resources: []*ast.Resource{
									{
										Path:   "/",
										Method: "GET",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			want: &ast.Cell{
				Token: token.Token{
					Type:    token.CELL,
					Literal: "Cell",
				},
				Name: &ast.Identifier{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "hr",
					},
					Value: "hr",
				},
				Components: []*ast.Component{
					{
						Token: token.Token{
							Type:    token.COMPONENT,
							Literal: "Component",
						},
						Name: &ast.Identifier{
							Token: token.Token{
								Type:    token.IDENTIFIER,
								Literal: "hr",
							},
							Value: "hr",
						},
						Image: &ast.Image{
							Type: "Docker",
							Name: "docker.io/wso2vick/sampleapp-hr",
						},
						Ports: []*ast.Port{
							{
								Type:          "TCP",
								ContainerPort: 8080,
								HostPort:      80,
							},
						},
						Env: []*ast.Env{
							{
								Key:   "employeegw_url",
								Value: "employee",
								Type:  "Cell",
							},
							{
								Key:   "stockgw_url",
								Value: "stock-options",
								Type:  "Cell",
							},
						},
					},
				},
				Ingresses: []*ast.Ingress{
					{
						Token: token.Token{
							Type:    token.INGRESS,
							Literal: "Ingress",
						},
						Type: "HTTP",
						Routes: []*ast.Route{
							{
								Context: "/info",
								Backend: &ast.Identifier{
									Token: token.Token{
										Type:    token.IDENTIFIER,
										Literal: "hr",
									},
									Value: "hr",
								},
								Resources: []*ast.Resource{
									{
										Path:   "/",
										Method: "GET",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	opt := cmpopts.IgnoreFields(token.Token{}, "Line", "Column")
	for i, test := range tests {
		got := file.Cells[i]
		if diff := cmp.Diff(test.want, got, opt); diff != "" {
			t.Errorf("AST is invalid (-want, +got)\n%v", diff)
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
