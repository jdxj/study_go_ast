package chapter2

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestBasicLit(t *testing.T) {
	var lit9527 = &ast.BasicLit{
		Kind:  token.INT,
		Value: "9527",
	}
	ast.Print(nil, lit9527)
}

func TestParser(t *testing.T) {
	// 返回一个语法树
	expr, _ := parser.ParseExpr(`9527`)
	ast.Print(nil, expr)

	expr, _ = parser.ParseExpr(`"9527"`)
	ast.Print(nil, expr)
}

func TestIdent(t *testing.T) {
	ast.Print(nil, ast.NewIdent(`x`))

	expr, _ := parser.ParseExpr(`x`)
	ast.Print(nil, expr)
}
