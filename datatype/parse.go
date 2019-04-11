package main

import "fmt"

type Test struct {
	x, y int
}

func main() {
	i := 1
	fmt.Println(i)

	handle(i)
	fmt.Println(i)
	handle1(&i)
	fmt.Println(i)

	t := Test{
		x: 3,
		y: 6,
	}
	handle2(t)
	fmt.Println(t)
	handle3(&t)
	fmt.Println(t)

}

func handle(i int) {
	i = i + 100
}

func handle1(i *int) {
	*i = *i + 100
}

func handle2(t Test) {
	t.x += 100
	t.y += 200
}
func handle3(t *Test) {
	t.x += 100
	t.y += 200
}
