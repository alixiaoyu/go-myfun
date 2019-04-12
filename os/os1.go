package main

import (
	"fmt"
	"os"
)

func main() {
	// io.WriteString()
	path, err := os.Getwd()
	if err != nil {
		return
	}
	envp := os.Getenv("GOBIN")

	fmt.Println(os.Chdir("/Users/lixiaoyu/go/src"))
	fmt.Println(os.Getwd())

	fmt.Println("当前目录是：", path)
	fmt.Println("当前环境是：", envp)
}
