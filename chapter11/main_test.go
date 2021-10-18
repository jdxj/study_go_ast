package main

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"testing"
)

func TestTypes(t *testing.T) {
	src := `
package pkg

func hello() {
	var _ = "a" + 1
}`
	fset := token.NewFileSet()
	// 得到语法树
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	// 语义检查
	/*
		第一个参数表示要检查包的路径
		第二个参数表示全部的文件集合（用于将语法树中元素的位置信息解析为文件名和行列号）
		第三个参数是该包中所有文件对应的语法树
		最后一个参数可用于存储检查过程中产生的分析结果

		如果成功该方法返回一个types.Package对象，表示当前包的信息。
	*/
	pkg, err := new(types.Config).Check("hello.go", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	_ = pkg
}

func TestType2(t *testing.T) {
	src := `package main

import "math"

func main() {
	var _ = "a" + math.Pi
}`
	fset := token.NewFileSet()
	// 得到语法树
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	// import "go/importer"
	// 处理导入其它包
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("hello.go", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}
	_ = pkg
}
