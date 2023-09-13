package main

import "fmt"

func test(fn func() int) int {
	return fn()
}
func fn() int {
	return 200
}
func main() {
	//这是直接使用匿名函数
	s1 := test(func() int { return 100 })
	fmt.Println(s1)
	//这是传入一个函数
	s1 = test(fn)
	fmt.Println(s1)
}
