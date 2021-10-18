package main

import (
	"go/ast"
	"go/parser"
	"testing"
)

func printAst(src string) {
	expr, _ := parser.ParseExpr(src)
	ast.Print(nil, expr)
}

func TestParse1(t *testing.T) {
	src := `int(x)`
	printAst(src)
}

func TestParse2(t *testing.T) {
	src := `x.y`
	printAst(src)
}

func TestParse3(t *testing.T) {
	src := `x[y]`
	printAst(src)
}

func TestParse4(t *testing.T) {
	src := `x[1:2:3]`
	printAst(src)
}

func TestParse5(t *testing.T) {
	src := `x.(y)`
	printAst(src)
}
