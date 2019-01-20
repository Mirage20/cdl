package vick

import (
	"fmt"
	"github.com/mirage20/cdl/ast"
	"github.com/mirage20/cdl/generator/vick/v1alpha1"
)

func CreateK8SCell(cell *ast.Cell) *v1alpha1.Cell {
	return &v1alpha1.Cell{
		TypeMeta: v1alpha1.TypeMeta{
			APIVersion: "vick.wso2.com/v1alpha1",
			Kind:       "Cell",
		},
		ObjectMeta: v1alpha1.ObjectMeta{
			Name: cell.Name.Value,
		},
		Spec: v1alpha1.CellSpec{
			GatewayTemplate: v1alpha1.GatewayTemplateSpec{
				Spec: v1alpha1.GatewaySpec{
					APIRoutes: createAPIRoutes(cell.Ingresses[0].Routes),
				},
			},
			ServiceTemplates: createServices(cell.Components),
		},
	}
}

func createAPIRoutes(routes []*ast.Route) []v1alpha1.APIRoute {
	var apiRoute []v1alpha1.APIRoute
	for _, route := range routes {
		var apiDefinitions []v1alpha1.APIDefinition
		for _, resource := range route.Resources {
			apiDefinitions = append(apiDefinitions, v1alpha1.APIDefinition{
				Method: resource.Method,
				Path:   resource.Path,
			})
		}
		apiRoute = append(apiRoute, v1alpha1.APIRoute{
			Context:     route.Context,
			Definitions: apiDefinitions,
			Backend:     route.Backend.Value,
			Global:      false,
		})
	}
	return apiRoute
}

func createServices(components []*ast.Component) []v1alpha1.ServiceTemplateSpec {
	var serviceSpec []v1alpha1.ServiceTemplateSpec
	var one int32 = 1
	for _, component := range components {
		serviceSpec = append(serviceSpec, v1alpha1.ServiceTemplateSpec{
			ObjectMeta: v1alpha1.ObjectMeta{
				Name: component.Name.Value,
			},
			Spec: v1alpha1.ServiceSpec{
				Replicas:    &one,
				ServicePort: int32(component.Ports[0].HostPort),
				Container: v1alpha1.Container{
					Image: component.Image.Name,
					Ports: createContainerPorts(component.Ports),
					Env:   createEnvVars(component.Env),
				},
			},
		})
	}

	serviceSpec = append(serviceSpec, v1alpha1.ServiceTemplateSpec{
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
	})
	return serviceSpec
}

func createEnvVars(envs []*ast.Env) []v1alpha1.EnvVar {
	var envVars []v1alpha1.EnvVar
	for _, env := range envs {
		var val string
		switch env.Type {
		case "Cell":
			val = fmt.Sprintf("%s--gateway-service", env.Value)
		default:
			val = env.Value
		}
		envVars = append(envVars, v1alpha1.EnvVar{
			Name:  env.Key,
			Value: val,
		})
	}
	return envVars
}

func createContainerPorts(ports []*ast.Port) []v1alpha1.ContainerPort {
	var cPorts []v1alpha1.ContainerPort
	for _, port := range ports {
		cPorts = append(cPorts, v1alpha1.ContainerPort{
			ContainerPort: int32(port.ContainerPort),
		})
	}
	return cPorts
}
