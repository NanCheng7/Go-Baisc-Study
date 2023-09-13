package main

import "fmt"

// 支持返回值 命名 ，默认值为类型零值,命名返回参数可看做与形参类似的局部变量,由return隐式返回
func f1() (names []string, m map[string]int, num int) {
	m = make(map[string]int)
	m["k1"] = 2

	return
}

func main() {
	a, b, c := f1()
	fmt.Println(a, b, c)
}
