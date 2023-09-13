package main

import "fmt"

// 多个返回值 用括号扩起来
func sum(a, b int) (int, int) {
	return a, b
}
func main() {
	a, b := sum(2, 3)
	fmt.Println(a, b)
}
