package main

import "fmt"

func main() {
	step := 2
	for step > 0 {
		step--
		fmt.Println(step)
		//执行一次就结束了
		return
	}
	//不会执行
	fmt.Println("结束之后的语句....")
}
