package main

import (
	"fmt"
)

func producer2(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square2(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()

	return out
}

func main() {
	in := producer2(1, 2, 3, 4)
	ch := square2(in)

	// consumer
	for ret := range ch {
		fmt.Printf("%3d", ret)
	}
	fmt.Println()
}
