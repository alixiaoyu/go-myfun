package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()
	go func() {
		for {
			x := <-ch1
			ch2 <- x * x
		}
	}()

	for {
		fmt.Println("----->", <-ch2)
	}
}
