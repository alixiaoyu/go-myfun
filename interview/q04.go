package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var num, char, wg sync.WaitGroup
	num.Add(1)
	wg.Add(2)

	go func() {
		for i := 1; i <= 10; i += 2 {
			char.Wait()
			char.Add(1)
			for j := i; j <= 10 && j < i+2; j++ {
				fmt.Print(j)
			}
			num.Done()
		}
		wg.Done()
	}()

	go func() {
		s := "ABCDEFGHIJ"
		for i := 0; i < len(s); i += 2 {
			num.Wait()
			num.Add(1)
			for j := i; j < len(s) && j < i+2; j++ {
				fmt.Print(s[j : j+1])
			}
			char.Done()
		}
		wg.Done()
	}()

	wg.Wait()
}
