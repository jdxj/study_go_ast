# 语句块和语句

语句块和语句只能在函数体内部定义.

# 语句

- 声明语句
- 标签语句
- 普通表达式语句
- 控制流语句

```
FunctionBody  = Block .

Block         = "{" StatementList "}" .
StatementList = { Statement ";" } .

Statement     = Declaration | LabeledStmt | SimpleStmt
              | GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt
              | FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt
              | DeferStmt
              .
```

Block 也是一种合法的语句 (注意文法 `Statement` 中又包含了 `Block`).

实际上定义空的语句块并不能算真正的语句, 它只是在编译阶段定义新的变量作用域, 并没有产生新的语句或计算.
最简单的语句是表达式语句，不管是简单表达式还是复杂的表达式都可以作为一个独立的语句。

```go
func main() {
	42 // 在文法角度是合法的, 在语义角度不合法: 未使用
}
```

表达式中可能还有函数调用, 而函数调用可能有其它的副作用, 因此表达式语句一般常用于触发函数调用.

## 声明语句

函数中除了输入参数和返回值参数之外, 还可以定义临时的局部变量保存函数的状态.
如果临时变量被闭包函数捕获, 那么临时变量维持的函数状态将伴随闭包函数的整个生命周期.

## for

四种形式:

```go
for {}
for true {}
for i := 0; true; i++ {}
for i, v := range m {}
```

## go 和 defer 语句

他两类似.