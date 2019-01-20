package vick

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/mirage20/cdl/ast"
	"strings"
)

func Generate(file *ast.File, output string) ([]byte, error) {
	var result []byte
	outFormat := strings.ToLower(output)
	for _, cell := range file.Cells {
		if outFormat == "yaml" {
			bytes, err := yaml.Marshal(CreateK8SCell(cell))
			if err != nil {
				return nil, err
			}
			result = append(result, bytes...)
			result = append(result, []byte("---\n")...)
		} else {
			bytes, err := json.MarshalIndent(CreateK8SCell(cell), "", "  ")
			if err != nil {
				return nil, err
			}
			result = append(result, bytes...)
			result = append(result, '\n')
		}
	}
	return result, nil
}
