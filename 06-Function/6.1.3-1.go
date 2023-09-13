package main

import (
	"fmt"
)

/* 定义相互交换值的函数 */
func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	var a, b int = 1, 2
	/*
	   调用 swap() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)

	fmt.Println(a, b)
}
