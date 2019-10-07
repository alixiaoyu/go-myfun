// Add函数一定要在Wait函数执行前执行
// package main

// import (
// 	"fmt"
// 	"log"
// 	"sync"
// )

// func father(wg *sync.WaitGroup) {
// 	wg.Add(1)
// 	defer wg.Done()

// 	fmt.Printf("father\n")
// 	for i := 0; i < 10; i++ {
// 		go child(wg, i)
// 	}
// }

// func child(wg *sync.WaitGroup, id int) {
// 	wg.Add(1)
// 	defer wg.Done()

// 	fmt.Printf("child [%d]\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	go father(&wg)

// 	wg.Wait()
// 	log.Printf("main: father and all chindren exit")
// 	return
// }

package main

import (
	"fmt"
	"log"
	"sync"
)

// Golang的WaitGroup陷阱
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go father(&wg)
	wg.Wait()
	log.Printf("main: father and all chindren exit")
}

func father(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("father\n")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go child(wg, i)
	}
}

func child(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("child [%d]\n", id)
}
