package main

import (
	"fmt"
	"github.com/mirage20/cdl/ast"
	"github.com/mirage20/cdl/generator"
	"github.com/mirage20/cdl/lexer"
	"github.com/mirage20/cdl/parser"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

var (
	runtime   string
	outFormat string
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "cdl FILE",
		Short:         "Generate runtime artifacts from Cell definition",
		SilenceUsage:  false,
		SilenceErrors: true,
		Version:       "0.1.0",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("please provide cdl file(s) as arguments")
			}
			if strings.ToLower(runtime) != "kubernetes" {
				return fmt.Errorf("runtime %q is not supported", runtime)
			}

			pFileCombined := &ast.File{}
			for _, filePath := range args {
				fileBytes, err := ioutil.ReadFile(filePath)
				if err != nil {
					return err
				}

				l := lexer.New(string(fileBytes))
				p := parser.New(l)
				pFile := p.ParseFile()
				checkParserErrors(p, filePath)
				if len(pFile.Cells) == 0 {
					return fmt.Errorf("file %q should contain at least 1 cell definition", filePath)
				}
				pFileCombined.Cells = append(pFileCombined.Cells, pFile.Cells...)
			}

			outBytes, err := generator.Generate(pFileCombined, generator.KUBERNETES, outFormat)
			if err != nil {
				return err
			}
			fmt.Println(string(outBytes))
			return nil
		},
	}
	cmd.Flags().StringVarP(&runtime, "runtime", "r", "kubernetes", "Target runtime for artifact generation.")
	cmd.Flags().StringVarP(&outFormat, "output", "o", "json", "Output format. One of json,yaml")
	return cmd
}

func main() {
	cmd := newRootCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Printf(fmt.Sprintf("%s: %s\n", "cdl", err))
		os.Exit(1)
	}
}

func checkParserErrors(p *parser.Parser, f string) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	fmt.Printf("Error while parsing file %q.\n", f)
	for _, msg := range errors {
		fmt.Printf("parser error: %q\n", msg)
	}
	os.Exit(1)
}
