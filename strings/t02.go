package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("--->", t.Format("2006-01-02"))
	fmt.Printf("%02d.%02d.%04d", t.Day(), t.Month(), t.Year())
}
