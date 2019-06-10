package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := 7001.836500000000000000000001
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", v), 3)
	fmt.Println(res)
}
