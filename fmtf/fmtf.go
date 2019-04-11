package main

import (
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	p := Point{
		X: 1,
		Y: 2,
	}
	fmt.Printf("%v\n", p)
	// 结构体%+v 格式化输出字段名
	fmt.Printf("%+v\n", p)
	// %#v 输出go' 的语法表示
	fmt.Printf("%#v\n", p)
	// %T
	fmt.Printf("%T\n", p)
	// %s
	fmt.Printf("%s\n", "\"string\"")
	// %q
	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
