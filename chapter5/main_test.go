package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func TestParseImport(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", srcImport, parser.ImportsOnly)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range f.Imports {
		fmt.Printf("import: name = %v, path = %#v\n", s.Name, s.Path)
	}
}

const srcImport = `package foo
import "pkg-a"
import pkg_b_v2 "pkg-b"
import . "pkg-c"
import _ "pkg-d"
`

func TestParseType(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", srcType, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				fmt.Printf("%T\n", spec)
			}
		}
	}
}

const srcType = `package foo
type MyInt1 int
type MyInt2 = int
`

func TestParseConst(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", srcConst, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				fmt.Printf("%T\n", spec)
				ast.Print(nil, spec)
			}
		}
	}
}

const srcConst = `package foo
const Pi = 3.14
const E float64 = 2.71828
const a, b = 1, 2`

func TestParseVal(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", srcVal, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			fmt.Printf("token: %v\n", v.Tok)
			fmt.Printf("Lparen: %d, Rparen: %d\n", v.Lparen, v.Rparen)
			for _, spec := range v.Specs {
				ast.Print(nil, spec)
			}
		}
	}
}

const srcVal = `package foo
var Pi = 3.14
`
