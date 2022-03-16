package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fs := token.NewFileSet()
	f, _ := parser.ParseFile(fs, "./example/example.go", nil, 0)

	for _, d := range f.Decls {
		ast.Print(fs, d)
	}
}
