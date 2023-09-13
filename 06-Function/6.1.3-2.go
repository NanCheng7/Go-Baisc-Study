package main

import (
	"fmt"
)

func test(s string, n ...int) string {
	var x int
	//对切片n进行累和
	for _, i := range n {
		x += i
	}
	//输出printf的格式模板
	return fmt.Sprintf(s, x)
}

func main() {
	s := []int{1, 2, 3}
	res := test("sum: %d", s...) // slice... 展开slice
	println(res)
}
