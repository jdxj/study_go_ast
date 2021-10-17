# 复合类型

无法用一个标识符表示的类型.

## 数组类型

`*ast.ArrayType`:

```go
type ArrayType struct {
	Lbrack token.Pos // position of "["
	Len    Expr      // Ellipsis node for [...]T array types, nil for slice types
	Elt    Expr      // element type
}
```

## 切片类型

也使用 `*ast.ArrayType` 表示, 但是 `Len` 字段为 `nil`.

## 函数类型

函数类型只包含函数签名:

- 参数
- 返回值
