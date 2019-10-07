package main

import (
	"fmt"
)

type Radix struct {
	length int     //序列中最大数的位数
	radix  [][]int //0-9的10个桶
	nums   []int   //要排序的序列
}

//初始化排序的相关参数
func (r *Radix) Init(numbers []int) {
	r.nums = numbers
	r.getLen()
}

//获取序列最大数的位数
func (r *Radix) getLen() {
	size := len(r.nums)
	max := r.nums[0]
	for i := 1; i < size; i++ {
		if r.nums[i] > max {
			max = r.nums[i]
		}
	}
	r.length = 1
	max = max / 10
	for max > 0 {
		r.length++
		max = max / 10
	}
}

//输出序列
func (r *Radix) Display() {
	for _, v := range r.nums {
		fmt.Printf("%d\t", v)
	}
	fmt.Print("\n")
}

//排序函数
func (r *Radix) Sort() {
	size := len(r.nums)
	var i, j, k int
	m := 1
	for i = 1; i <= r.length; i++ { //循环次数为常数，即序列中最大数的位数
		r.radix = make([][]int, 10)
		for j = 0; j < 10; j++ { //初始化0-9各位桶
			r.radix[j] = make([]int, 0)
		}
		for j = 0; j < size; j++ { //遍历要排序的序列，将相应位数的元素加入到对应的位桶中
			k = r.nums[j] / m % 10
			r.radix[k] = append(r.radix[k], r.nums[j])
		}
		r.nums = make([]int, 0)  //清空原序列数组
		for j = 9; j >= 0; j-- { //将9-0各位桶的数重新组装放到原序列数组当中
			r.nums = append(r.nums, r.radix[j]...)
		}
		m = m * 10
		r.Display()
	}

}

func main() {
	r := new(Radix)
	r.Init([]int{430, 122, 332, 167, 899, 998, 455, 691, 571})
	r.Sort()
}
