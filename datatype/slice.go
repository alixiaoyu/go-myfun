package main

import (
	"fmt"
)

/* 一个slice是一个轻量级的数据结构，
提供了访问数组子序列（或者全部）元素的功能，
而且slice的底层确实引用一个数组对象。
一个slice由三个部分构成：指针、长度和容量。
指针指向第一个slice元素对应的底层数组元素的地址，
要注意的是slice的第一个元素并不一定就是数组的第一个元素。
长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。
内置的len和cap函数分别返回slice的长度和容量 */

var month = [...]string{1: "January", 2: "February", 3: "March"}

func main() {
	s := []int{0, 45, 34, 3, 4, 5, 82, 7}
	reverse(s)
	fmt.Println("--->", s)
	x := []int{}
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...)
	fmt.Println(x)

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", noEmpty(data))
	fmt.Printf("%q\n", data)
}

// 反转数组
func reverse(s []int) {
	l := len(s) - 1
	for i := 0; i <= l/2; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}

}

// 返回不包含空格
func noEmpty(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}
