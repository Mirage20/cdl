package generator

import (
	"errors"
	"github.com/mirage20/cdl/ast"
	"github.com/mirage20/cdl/generator/vick"
)

type Generator int

const (
	KUBERNETES Generator = iota
)

func Generate(file *ast.File, generator Generator, output string) ([]byte, error) {
	switch generator {
	case KUBERNETES:
		return vick.Generate(file, output)
	default:
		return nil, errors.New("unknown generator")
	}
}
