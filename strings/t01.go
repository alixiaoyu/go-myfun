package main

import (
	"fmt"
	"strings"
)

func main() {
	isHas := strings.HasPrefix("string", "str")
	isCon := strings.Contains("strings", "ings")
	in := strings.Index("strings", "g")
	sp := strings.Replace("aaabbbcccaaadddaaa", "aaa", "AAA", -1)
	sa := strings.ReplaceAll("aaabbbcccaaadddaaa", "aaa", "EEE")

	fmt.Println("---->", isHas, isCon, in, sp, sa)
}
