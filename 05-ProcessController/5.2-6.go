package main

import "fmt"

func main() {
	step := 2
	for step > 0 {
		step--
		fmt.Println(step)
		panic("出错了")
	}
	// 不会执行
	fmt.Println("结束之后的语句。。。")
}
