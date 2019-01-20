package lexer

import (
	"testing"

	"github.com/mirage20/cdl/token"
)

func TestNextTokenOneCell(t *testing.T) {
	input := `

       
    # This is a comment 
 		 
    # This is another comment
    
    Cell employee {
      
    	Component employee {
            Image:Docker = "docker.io/wso2vick/sampleapp-employee"
            Ports {
				TCP 80->8080
				TCP 443->8081
			}
        }
   
        Component salary-service {
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
    	
	
    
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CELL, "Cell"},
		{token.IDENTIFIER, "employee"},
		{token.LBRACE, "{"},
		{token.COMPONENT, "Component"},
		{token.IDENTIFIER, "employee"},
		{token.LBRACE, "{"},
		{token.IMAGE, "Image"},
		{token.COLON, ":"},
		{token.DOCKER, "Docker"},
		{token.ASSIGN, "="},
		{token.STRING, "docker.io/wso2vick/sampleapp-employee"},
		{token.PORTS, "Ports"},
		{token.LBRACE, "{"},
		{token.TCP, "TCP"},
		{token.NUMBER, "80"},
		{token.RARROW, "->"},
		{token.NUMBER, "8080"},
		{token.TCP, "TCP"},
		{token.NUMBER, "443"},
		{token.RARROW, "->"},
		{token.NUMBER, "8081"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.COMPONENT, "Component"},
		{token.IDENTIFIER, "salary-service"},
		{token.LBRACE, "{"},
		{token.IMAGE, "Image"},
		{token.COLON, ":"},
		{token.DOCKER, "Docker"},
		{token.ASSIGN, "="},
		{token.STRING, "docker.io/wso2vick/sampleapp-salary"},
		{token.PORTS, "Ports"},
		{token.LBRACE, "{"},
		{token.TCP, "TCP"},
		{token.NUMBER, "80"},
		{token.RARROW, "->"},
		{token.NUMBER, "8080"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.INGRESS, "Ingress"},
		{token.COLON, ":"},
		{token.HTTP, "HTTP"},
		{token.LBRACE, "{"},
		{token.STRING, "/employee"},
		{token.RARROW, "->"},
		{token.IDENTIFIER, "employee"},
		{token.LBRACE, "{"},
		{token.GET, "GET"},
		{token.STRING, "/"},
		{token.POST, "POST"},
		{token.STRING, "/"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenTwoCells(t *testing.T) {
	input := `

       
    # This is a comment 


	Cell stock-options {
		Component stock {
			Ports {
				 TCP 80->8080
			}
			Image:Docker = "docker.io/wso2vick/sampleapp-stock"
		}
	
		Ingress:HTTP {
			"/stock" -> stock {
				GET "/"
			}
		}
	}
	
	# This is another comment
	
	Cell hr {
		Component hr {
			Image:Docker = "docker.io/wso2vick/sampleapp-hr"    # This is different comment
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
	
	# This is last comment
	
    	
	
    
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CELL, "Cell"},
		{token.IDENTIFIER, "stock-options"},
		{token.LBRACE, "{"},
		{token.COMPONENT, "Component"},
		{token.IDENTIFIER, "stock"},
		{token.LBRACE, "{"},
		{token.PORTS, "Ports"},
		{token.LBRACE, "{"},
		{token.TCP, "TCP"},
		{token.NUMBER, "80"},
		{token.RARROW, "->"},
		{token.NUMBER, "8080"},
		{token.RBRACE, "}"},
		{token.IMAGE, "Image"},
		{token.COLON, ":"},
		{token.DOCKER, "Docker"},
		{token.ASSIGN, "="},
		{token.STRING, "docker.io/wso2vick/sampleapp-stock"},
		{token.RBRACE, "}"},
		{token.INGRESS, "Ingress"},
		{token.COLON, ":"},
		{token.HTTP, "HTTP"},
		{token.LBRACE, "{"},
		{token.STRING, "/stock"},
		{token.RARROW, "->"},
		{token.IDENTIFIER, "stock"},
		{token.LBRACE, "{"},
		{token.GET, "GET"},
		{token.STRING, "/"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},

		{token.CELL, "Cell"},
		{token.IDENTIFIER, "hr"},
		{token.LBRACE, "{"},
		{token.COMPONENT, "Component"},
		{token.IDENTIFIER, "hr"},
		{token.LBRACE, "{"},
		{token.IMAGE, "Image"},
		{token.COLON, ":"},
		{token.DOCKER, "Docker"},
		{token.ASSIGN, "="},
		{token.STRING, "docker.io/wso2vick/sampleapp-hr"},
		{token.PORTS, "Ports"},
		{token.LBRACE, "{"},
		{token.TCP, "TCP"},
		{token.NUMBER, "80"},
		{token.RARROW, "->"},
		{token.NUMBER, "8080"},
		{token.RBRACE, "}"},
		{token.ENV, "Env"},
		{token.LBRACE, "{"},
		{token.STRING, "employeegw_url"},
		{token.COLON, ":"},
		{token.CELL, "Cell"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "employee"},
		{token.STRING, "stockgw_url"},
		{token.COLON, ":"},
		{token.CELL, "Cell"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "stock-options"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.INGRESS, "Ingress"},
		{token.COLON, ":"},
		{token.HTTP, "HTTP"},
		{token.LBRACE, "{"},
		{token.STRING, "/info"},
		{token.RARROW, "->"},
		{token.IDENTIFIER, "hr"},
		{token.LBRACE, "{"},
		{token.GET, "GET"},
		{token.STRING, "/"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
