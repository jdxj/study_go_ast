package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func printBlock(src string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}
func TestParse1(t *testing.T) {
	src := `package pkgname
func main() {}`
	printBlock(src)
}

func TestParse2(t *testing.T) {
	src := `package pkgname
func main() {
	{}
	{}
}`
	printBlock(src)
}

func TestParse3(t *testing.T) {
	src := `package pkgname
func main() {
	42
}`
	printBlock(src)
}

func TestParse4(t *testing.T) {
	src := `package foo
func main() {
	return 42, err
}`
	printBlock(src)
}

func TestParse5(t *testing.T) {
	src := `package foo
func main() {
	var a int
}`
	printBlock(src)
}

func TestParse6(t *testing.T) {
	src := `package foo
func main() {
	a, b := 1, 2
}`
	printBlock(src)
}

func TestParse7(t *testing.T) {
	src := `package foo
func main() {
	if true {} else {}
}`
	printBlock(src)
}

func TestParse8(t *testing.T) {
	src := `package foo
func main() {
	for x; y; z {}
}`
	printBlock(src)
}

func a() {
	var b []int
	for range b {
	}
}

func TestParse9(t *testing.T) {
	src := `package foo
func main() {
	for range ch {}
}`
	printBlock(src)
}

func TestParse10(t *testing.T) {
	src := `package foo
func main() {
	x.(int)
}`
	printBlock(src)
}

func TestParse11(t *testing.T) {
	src := `package foo
func main() {
	go hello("光谷码农")
}`
	printBlock(src)
}
