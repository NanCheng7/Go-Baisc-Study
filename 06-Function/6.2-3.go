package main

import (
	"fmt"
)

// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	var sum int
	for _, v := range list {
		sum += v
		fmt.Println("sum: ", sum)
		f(v)
	}
}
func main() {
	// 使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println("single: ", v)
	})
}
