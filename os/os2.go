package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Stat("os1.go")
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println(file.Mode())

	err = os.Chmod("os1.go", 0777)
	if err != nil {
		fmt.Println("ERROR")
	}
	filemode, _ := os.Stat("os1.go")
	fmt.Println(filemode.Mode())
}
