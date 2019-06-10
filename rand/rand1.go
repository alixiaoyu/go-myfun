package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		// v := 4.635 + rand.Float64()*float64(rand.Intn(6)-3)
		// res, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", v), 64)
		res := 0.1512 * (float64(i) + rand.Float64())
		fmt.Println("----", res)
	}

}
