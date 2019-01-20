package generator

import (
	"github.com/google/go-cmp/cmp"
	"github.com/mirage20/cdl/lexer"
	"github.com/mirage20/cdl/parser"
	"io/ioutil"
	"testing"
)

func TestK8SCellGeneration(t *testing.T) {
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
	p := parser.New(l)
	file := p.ParseFile()
	checkParserErrors(t, p)
	if file == nil {
		t.Fatalf("ParseFile() returned nil")
	}

	bytesGot, err := Generate(file, VICK, "json")
	if err != nil {
		t.Fatal(err)
	}
	bytesWant, err := ioutil.ReadFile("output.json")
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(bytesWant), string(bytesGot)); diff != "" {
		t.Errorf("(-want, +got)\n%v", diff)
	}
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
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
