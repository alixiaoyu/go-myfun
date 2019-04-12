package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	err := os.Chtimes("os3.go", time.Now(), time.Now())
	if err != nil {
		fmt.Println("----->")
		return
	}
	fi, _ := os.Stat("os3.go")
	fmt.Println(fi.ModTime())

	// data := os.Environ()
	// fmt.Println(data)

	mapping := func(s string) string {
		m := map[string]string{"addr": "www.baidu.com", "xiaowei": "widuu"}
		return m[s]
	}
	data := "hello $xiaowei blog address $addr"
	fmt.Printf("%s\n", os.Expand(data, mapping))

	ess := "GOBIN PATH $GOBIN"
	fmt.Println(os.Expand(ess, os.Getenv))
	fmt.Println(os.ExpandEnv(ess))

	name, _ := os.Hostname()
	fmt.Println(name)

}
