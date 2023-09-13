package main

import "fmt"

// 定义函数类型。
type FuncType func(s string, x, y int) string

func format(fn FuncType, s string, x, y int) string {
	return fn(s, x, y)
}
func formatFun(s string, x, y int) string {
	return fmt.Sprintf(s, x, y)
}
func main() {
	s2 := format(formatFun, "%d, %d", 10, 20)
	fmt.Println(s2)
}
