package main

import (
	"fmt"
)

func main() {
	// var msgs []func()
	// var ss []string
	array := []string{"1", "2", "3", "4"}

	for _, e := range array {
		// ss = append(ss, e)
		// msgs = append(msgs, func() {
		fmt.Println(e)
		// })
	}
	// fmt.Println("===", ss)
	// for _, v := range msgs {
	// 	v()
	// }
}
