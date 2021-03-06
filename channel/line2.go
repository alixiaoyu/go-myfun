package main

import (
	"fmt"
	"sync"
)

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

func square(ss string, inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			fmt.Println("---", ss, n*n)
			out <- n * n
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
	// FAN-IN
	for _, c := range cs {
		go collect(c)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := producer(1, 2, 3, 4)

	// FAN-OUT
	c1 := square("c1", in)
	c2 := square("c2", in)
	c3 := square("c3", in)
	// consumer
	for ret := range merge(c1, c2, c3) {
		fmt.Printf("%3d ", ret)
	}
	fmt.Println()
}
