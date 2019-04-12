package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now().Unix()
	fmt.Println("=====>START")
	s := fibnacia(45)
	fmt.Println("======>END")
	t2 := time.Now().Unix()
	fmt.Printf("一共用时：%d秒，结果是%d \n", t2-t1, s)

}

func fibnacia(x int) int {
	if x < 2 {
		return 1
	}
	return fibnacia(x-1) + fibnacia(x-2)
}
