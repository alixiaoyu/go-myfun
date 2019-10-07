package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ary := []int{2, 5, 3, 19, 26, 4, 49, 38, 50, 17}
	fmt.Println(ary)
	bubble(ary)
}

func bubble(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Sorted array is: ", arr)
}

func quick_sort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	middle_part := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}

	low_part = quick_sort(low_part)
	high_part = quick_sort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}

func selectSort(arr []int) {
	var min int = 0
	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	fmt.Println("Sorted array:    ", arr)
}

func insert(arr []int) {
	var i, j int
	for i = 1; i < len(arr); i++ {
		for j = 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println("Sorted array is: ", arr)
}

func hell(arr []int) {
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}

	fmt.Println("Sorted array is: ", arr)
}
