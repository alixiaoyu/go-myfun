package main

import "fmt"

// Golang中我们使用Channel或者sync.Mutex等锁保护数据，有没有一种机制可以检测代码中的数据竞争呢？
// 使用race检查竟态

func main() {
	i := 0
	go func() {
		i++
	}()
	fmt.Println(i)
}
