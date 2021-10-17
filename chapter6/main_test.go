package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func TestParseFunc(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			ast.Print(nil, fn)
		}
	}
}

const src = `package abc
func (p *xType) Hello(arg1, arg2 int) (bool, error) {}`
