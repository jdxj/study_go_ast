package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)

func TestBinaryExpr(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+2*3`)
	ast.Print(nil, expr)
}

func TestEval(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+2*3`)
	fmt.Println(Eval(expr))
}

func Eval(exp ast.Expr) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	}
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) float64 {
	switch exp.Op {
	case token.ADD:
		return Eval(exp.X) + Eval(exp.Y)
	case token.MUL:
		return Eval(exp.X) * Eval(exp.Y)
	}
	return 0
}

func TestVar(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+2*3+x`)
	fmt.Println(Eval2(expr, map[string]float64{
		"x": 100,
	}))
}

func Eval2(exp ast.Expr, vars map[string]float64) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr2(exp, vars)
	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	case *ast.Ident:
		return vars[exp.Name]
	}
	return 0
}

func EvalBinaryExpr2(exp *ast.BinaryExpr, vars map[string]float64) float64 {
	switch exp.Op {
	case token.ADD:
		return Eval2(exp.X, vars) + Eval2(exp.Y, vars)
	case token.MUL:
		return Eval2(exp.X, vars) * Eval2(exp.Y, vars)
	}
	return 0
}
