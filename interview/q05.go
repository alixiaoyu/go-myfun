package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(in, quit <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			fmt.Println("收到退出信号")
			return // 必须return，否则goroutine是不会结束的
		case v := <-in:
			fmt.Println(v)
		}
	}
}

func main() {
	quit := make(chan int) // 退出通道
	in := make(chan int)

	wg.Add(100)
	go worker(in, quit)

	for i := 0; i < 3; i++ {
		in <- i
		time.Sleep(1 * time.Second)
	}
	// quit <- 1 // 向quit通道发送退出信号
	close(quit)
	wg.Wait()
}
