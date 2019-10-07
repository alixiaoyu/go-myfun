package main

import "fmt"

func main() {
	in := producer(1, 2, 3, 4)
	out := square(in)
	for i := range out {
		fmt.Println("-->", i)
	}
	fmt.Println()
}

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- i * i
		}
	}()
	return out
}
