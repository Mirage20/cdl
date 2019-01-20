package v1alpha1

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestCellType(t *testing.T) {
	var one int32 = 1
	cell := &Cell{
		TypeMeta: TypeMeta{
			APIVersion: "vick.wso2.com/v1alpha1",
			Kind:       "Cell",
		},
		ObjectMeta: ObjectMeta{
			Name: "employee",
		},
		Spec: CellSpec{
			GatewayTemplate: GatewayTemplateSpec{
				Spec: GatewaySpec{
					APIRoutes: []APIRoute{
						{
							Context: "employee",
							Definitions: []APIDefinition{
								{
									Method: "GET",
									Path:   "/",
								},
							},
							Backend: "employee",
							Global:  false,
						},
					},
				},
			},
			ServiceTemplates: []ServiceTemplateSpec{
				{
					ObjectMeta: ObjectMeta{
						Name: "employee",
					},
					Spec: ServiceSpec{
						Replicas:    &one,
						ServicePort: 80,
						Container: Container{
							Image: "docker.io/wso2vick/sampleapp-employee",
						},
					},
				},
				{
					ObjectMeta: ObjectMeta{
						Name: "salary",
					},
					Spec: ServiceSpec{
						Replicas:    &one,
						ServicePort: 80,
						Container: Container{
							Image: "docker.io/wso2vick/sampleapp-salary",
						},
					},
				},
				{
					ObjectMeta: ObjectMeta{
						Name: "debug",
					},
					Spec: ServiceSpec{
						Replicas:    &one,
						ServicePort: 80,
						Container: Container{
							Image: "docker.io/mirage20/k8s-debug-tools",
						},
					},
				},
			},
		},
	}

	bytes, err := json.MarshalIndent(cell, "", "  ")

	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("output.json", bytes, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
