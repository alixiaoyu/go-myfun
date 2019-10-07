package main

import (
	"fmt"
	"sync"
)

func main() {
	in := producer3(1, 2, 3, 4)

	// FAN-OUT
	c1 := square3(in)
	c2 := square3(in)
	c3 := square3(in)

	// wg := sync.WaitGroup{}
	// wg.Add(3)
	// out := merge3(&wg, c1, c2, c3)
	out := mergesss(c1, c2, c3)
	for i := range out {
		fmt.Println("--->", i)
	}
	// wg.Wait()

}

func producer3(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square3(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- i * i
		}
	}()
	return out
}

func mergesss(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))

	// ?---->
	// collect := func(in <-chan int) {
	// 	defer wg.Done()
	// 	for n := range in {
	// 		out <- n
	// 	}
	// }
	// for _, c := range ins {
	// 	go collect(c)
	// }

	for _, in := range ins {
		go func() {
			defer wg.Done()
			for i := range in {
				out <- i
			}

		}()
	}
	wg.Wait()
	close(out)
	// go func() {
	// 	wg.Wait()
	// 	close(out)
	// }()
	return out

}

func merge3(cs ...<-chan int) <-chan int {
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
