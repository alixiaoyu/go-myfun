package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().AddDate(0, -4, -27).Format("01-02")
	year := time.Now().Format("2006")
	fmt.Println("---->", t, year)
}
