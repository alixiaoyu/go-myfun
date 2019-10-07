package main

import "fmt"

func main() {
	fmt.Println("test1")
	fmt.Printf("in main: x = %d\n", f1())
	fmt.Println("test2")
	fmt.Printf("in main: x = %d\n", f2())
	fmt.Println("test3")
	fmt.Printf("in main: x = %d\n", f3())
	// fmt.Println("test4")
	// fmt.Printf("in main: x = %d\n", f4())
}

func f1() (result int) {
	defer func() {
		result++
		fmt.Printf("in defer: x = %d\n", result)
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Printf("in defer: x = %d\n", t)
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
		fmt.Printf("in defer: x = %d\n", r)
	}(r)
	return 1
}
