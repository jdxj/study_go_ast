package main

import (
	"go/ast"
	"go/parser"
	"testing"
)

func TestParseFuncLit(t *testing.T) {
	expr, _ := parser.ParseExpr(`func(){}`)
	ast.Print(nil, expr)
}

func TestParseArrayLit(t *testing.T) {
	expr, _ := parser.ParseExpr(`[...]int{1,2:3}`)
	ast.Print(nil, expr)
}

func TestParseStructLit(t *testing.T) {
	expr, _ := parser.ParseExpr(`struct{X int}{X:1}`)
	ast.Print(nil, expr)

	expr, _ = parser.ParseExpr(`struct{X int}{1}`)
	ast.Print(nil, expr)
}

func TestParseMapLit(t *testing.T) {
	expr, _ := parser.ParseExpr(`map[int]int{1:2}`)
	ast.Print(nil, expr)
}
