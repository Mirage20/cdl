package vick

import (
	"github.com/google/go-cmp/cmp"
	"github.com/mirage20/cdl/generator/vick/v1alpha1"
	"testing"

	"github.com/mirage20/cdl/lexer"
	"github.com/mirage20/cdl/parser"
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
	var one int32 = 1
	tests := []struct {
		want *v1alpha1.Cell
	}{
		{
			want: &v1alpha1.Cell{
				TypeMeta: v1alpha1.TypeMeta{
					APIVersion: "vick.wso2.com/v1alpha1",
					Kind:       "Cell",
				},
				ObjectMeta: v1alpha1.ObjectMeta{
					Name: "employee",
				},
				Spec: v1alpha1.CellSpec{
					GatewayTemplate: v1alpha1.GatewayTemplateSpec{
						Spec: v1alpha1.GatewaySpec{
							APIRoutes: []v1alpha1.APIRoute{
								{
									Context: "/employee",
									Definitions: []v1alpha1.APIDefinition{
										{
											Method: "GET",
											Path:   "/",
										},
										{
											Method: "POST",
											Path:   "/",
										},
									},
									Backend: "employee",
									Global:  false,
								},
							},
						},
					},
					ServiceTemplates: []v1alpha1.ServiceTemplateSpec{
						{
							ObjectMeta: v1alpha1.ObjectMeta{
								Name: "employee",
							},
							Spec: v1alpha1.ServiceSpec{
								Replicas:    &one,
								ServicePort: 80,
								Container: v1alpha1.Container{
									Image: "docker.io/wso2vick/sampleapp-employee",
									Ports: []v1alpha1.ContainerPort{
										{
											ContainerPort: 8080,
										},
										{
											ContainerPort: 8081,
										},
									},
								},
							},
						},
						{
							ObjectMeta: v1alpha1.ObjectMeta{
								Name: "salary",
							},
							Spec: v1alpha1.ServiceSpec{
								Replicas:    &one,
								ServicePort: 80,
								Container: v1alpha1.Container{
									Image: "docker.io/wso2vick/sampleapp-salary",
									Ports: []v1alpha1.ContainerPort{
										{
											ContainerPort: 8080,
										},
									},
								},
							},
						},
						{
							ObjectMeta: v1alpha1.ObjectMeta{
								Name: "debug",
							},
							Spec: v1alpha1.ServiceSpec{
								Replicas:    &one,
								ServicePort: 80,
								Container: v1alpha1.Container{
									Image: "docker.io/mirage20/k8s-debug-tools",
								},
							},
						},
					},
				},
			},
		},
		{
			want: &v1alpha1.Cell{
				TypeMeta: v1alpha1.TypeMeta{
					APIVersion: "vick.wso2.com/v1alpha1",
					Kind:       "Cell",
				},
				ObjectMeta: v1alpha1.ObjectMeta{
					Name: "stock-options",
				},
				Spec: v1alpha1.CellSpec{
					GatewayTemplate: v1alpha1.GatewayTemplateSpec{
						Spec: v1alpha1.GatewaySpec{
							APIRoutes: []v1alpha1.APIRoute{
								{
									Context: "/stock",
									Definitions: []v1alpha1.APIDefinition{
										{
											Method: "GET",
											Path:   "/",
										},
									},
									Backend: "stock",
									Global:  false,
								},
							},
						},
					},
					ServiceTemplates: []v1alpha1.ServiceTemplateSpec{
						{
							ObjectMeta: v1alpha1.ObjectMeta{
								Name: "stock",
							},
							Spec: v1alpha1.ServiceSpec{
								Replicas:    &one,
								ServicePort: 80,
								Container: v1alpha1.Container{
									Image: "docker.io/wso2vick/sampleapp-stock",
									Ports: []v1alpha1.ContainerPort{
										{
											ContainerPort: 8080,
										},
									},
								},
							},
						},
						{
							ObjectMeta: v1alpha1.ObjectMeta{
								Name: "debug",
							},
							Spec: v1alpha1.ServiceSpec{
								Replicas:    &one,
								ServicePort: 80,
								Container: v1alpha1.Container{
									Image: "docker.io/mirage20/k8s-debug-tools",
								},
							},
						},
					},
				},
			},
		},
		{
			want: &v1alpha1.Cell{
				TypeMeta: v1alpha1.TypeMeta{
					APIVersion: "vick.wso2.com/v1alpha1",
					Kind:       "Cell",
				},
				ObjectMeta: v1alpha1.ObjectMeta{
					Name: "hr",
				},
				Spec: v1alpha1.CellSpec{
					GatewayTemplate: v1alpha1.GatewayTemplateSpec{
						Spec: v1alpha1.GatewaySpec{
							APIRoutes: []v1alpha1.APIRoute{
								{
									Context: "/info",
									Definitions: []v1alpha1.APIDefinition{
										{
											Method: "GET",
											Path:   "/",
										},
									},
									Backend: "hr",
									Global:  false,
								},
							},
						},
					},
					ServiceTemplates: []v1alpha1.ServiceTemplateSpec{
						{
							ObjectMeta: v1alpha1.ObjectMeta{
								Name: "hr",
							},
							Spec: v1alpha1.ServiceSpec{
								Replicas:    &one,
								ServicePort: 80,
								Container: v1alpha1.Container{
									Image: "docker.io/wso2vick/sampleapp-hr",
									Ports: []v1alpha1.ContainerPort{
										{
											ContainerPort: 8080,
										},
									},
									Env: []v1alpha1.EnvVar{
										{
											Name:  "employeegw_url",
											Value: "employee--gateway-service",
										},
										{
											Name:  "stockgw_url",
											Value: "stock-options--gateway-service",
										},
									},
								},
							},
						},
						{
							ObjectMeta: v1alpha1.ObjectMeta{
								Name: "debug",
							},
							Spec: v1alpha1.ServiceSpec{
								Replicas:    &one,
								ServicePort: 80,
								Container: v1alpha1.Container{
									Image: "docker.io/mirage20/k8s-debug-tools",
								},
							},
						},
					},
				},
			},
		},
	}

	for i, test := range tests {
		got := CreateK8SCell(file.Cells[i])
		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("K8S Vick cell generation is invalid (-want, +got)\n%v", diff)
		}
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
