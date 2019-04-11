package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1.33333
	aa := "1.23445"
	s := strconv.FormatFloat(a, 'f', 3, 64)
	fmt.Printf("%#v\n", s)
	ss, _ := strconv.ParseFloat(aa, 64)
	fmt.Printf("%#v\n", ss)
}
