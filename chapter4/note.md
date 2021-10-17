# Go 语言代码结构层面

- 目录结构
- 目录内部的包结构
- 文件内部的代码结构

# 目录结构

- 同一包的测试文件, 可以属于独立的包 (原包名+`_test`)

# 文件结构

## 顶级语法元素

- package
- import
- type
- const
- var
- func

`*ast.GenDecl` 类型包括:

- import
- type
- const
- var

`*ast.FuncDecl` 类型包括:

- func