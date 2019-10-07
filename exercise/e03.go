// defer在defer语句处执行，defer的执行结果是把defer后的函数压入到栈，等待return或者函数panic后，再按先进后出的顺序执行被defer的函数。
// defer的函数的参数是在执行defer时计算的，defer的函数中的变量的值是在函数执行时计算的。
package main

import "fmt"

func main() {
	fmt.Println("test1")
	fmt.Printf("in main: x = %d\n", test1())
	fmt.Println("test2")
	fmt.Printf("in main: x = %d\n", test2())
	fmt.Println("test3")
	fmt.Printf("in main: x = %d\n", test3())
	fmt.Println("test4")
	fmt.Printf("in main: x = %d\n", test4())
}

func test1() (x int) {
	defer fmt.Printf("in defer: x = %d\n", x)
	x = 7
	return 9
}

func test2() (x int) {
	x = 7
	defer fmt.Printf("in defer: x = %d\n", x)
	return 9
}

func test3() (x int) {
	x = 7
	defer func(n int) {
		fmt.Printf("in defer: x = %d\n", n)
	}(x)

	// x = 7
	return 9
}

func test4() (x int) {
	defer func(n int) {
		fmt.Printf("in defer x as parameter: x = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", x)
	}(x)

	x = 7
	return 9
}
