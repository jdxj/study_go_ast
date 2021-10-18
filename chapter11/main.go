package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

type Program struct {
	fs   map[string]string         // 每个包对应的源代码字符串
	ast  map[string]*ast.File      // 每个包对应的语法树
	pkgs map[string]*types.Package // 经过语义检查的包对象
	fset *token.FileSet            // 文件的位置信息
}

func NewProgram(fs map[string]string) *Program {
	return &Program{
		fs:   fs,
		ast:  make(map[string]*ast.File),
		pkgs: make(map[string]*types.Package),
		fset: token.NewFileSet(),
	}
}

func (p *Program) LoadPackage(path string) (pkg *types.Package, f *ast.File, err error) {
	if pkg, ok := p.pkgs[path]; ok {
		return pkg, p.ast[path], nil
	}

	f, err = parser.ParseFile(p.fset, path, p.fs[path], parser.AllErrors)
	if err != nil {
		return nil, nil, err
	}

	/*
		因为没有初始化types.Config的Importer成员,
		因此目前该方法只能加载没有导入其他包的叶子类型的包（对应math包就是这种类型）.
		比如叶子类型的math包被加载成功之后, 则会被记录到Program对象的ast和pkgs成员中.
		然后当遇到已经被记录过的叶子包被导入时，就可以复用这些信息.
	*/
	//conf := types.Config{Importer: nil}

	/*
		Program类型实现了types.Importer接口, 就可以用于types.Config的包加载工作.
	*/
	conf := types.Config{Importer: p}
	pkg, err = conf.Check(path, p.fset, []*ast.File{f}, nil)
	if err != nil {
		return nil, nil, err
	}

	p.ast[path] = f
	p.pkgs[path] = pkg
	return pkg, f, nil
}

func (p *Program) Import(path string) (*types.Package, error) {
	if pkg, ok := p.pkgs[path]; ok {
		return pkg, nil
	}
	//return nil, fmt.Errorf("not found: %s", path)

	// 递归处理, 遇到没有的 path 要重新从文件读取
	pkg, _, err := p.LoadPackage(path)
	return pkg, err
}

func main() {
	prog := NewProgram(map[string]string{
		"hello": `
			package main
			import "math"
			func main() { var _ = 2 * math.Pi }
		`,
		// 模拟 math 包
		"math": `
			package math
			const Pi = 3.1415926
		`,
	})

	_, _, err := prog.LoadPackage("math")
	if err != nil {
		log.Fatal(err)
	}

	pkg, f, err := prog.LoadPackage("hello")
	if err != nil {
		log.Fatal(err)
	}
	_ = pkg
	_ = f
}
