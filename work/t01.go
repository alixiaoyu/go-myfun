package main

import (
	"fmt"
	"strings"
)

func main() {
	videoIds := "151,152,153,154,155,156,157,158,159,160,161,162,163,164"

	s := strings.Split(videoIds, ",")
	fmt.Println("--->", s, len(s))

	println("(" + strings.Join(s, ",") + ")")
}
