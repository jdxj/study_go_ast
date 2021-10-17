package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func print(src string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	for _, decl := range f.Decls {
		ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
	}
}

const src1 = `package foo
type Int1 int
type Int2 pkg.int
`

func TestParse1(t *testing.T) {
	print(src1)
}

const src2 = `package foo
type IntPtr *int
`

const src3 = `package foo
type IntPtrPtr **int
`

func TestParse2And3(t *testing.T) {
	print(src2)
	print(src3)
}

const src4 = `package foo
type IntArray [1]int`

const src5 = `package foo
type IntArrayArray [1][2]int`

func TestParse4And5(t *testing.T) {
	print(src4)
	print(src5)
}

const src6 = `package foo
type IntSlice []int`

func TestParse6(t *testing.T) {
	print(src6)
}

const src7 = `package foo
type MyStruct struct {
	a, b int "int value"
	string
}`

func TestParse7(t *testing.T) {
	print(src7)
}

const src8 = `package foo
type IntStringMap map[int]string`

func TestParse8(t *testing.T) {
	print(src8)
}

const src9 = `package foo
type IntChan chan int`

func TestParse9(t *testing.T) {
	print(src9)
}

const src10 = `package foo
type IntReader interface {
	Read() int
}`

func TestParse10(t *testing.T) {
	print(src10)
}
