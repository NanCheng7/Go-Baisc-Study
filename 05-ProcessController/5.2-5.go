package main

import "fmt"

func main() {
	step := 2
	for step > 0 {
		step--
		fmt.Println(step)
		//跳出循环,还会继续执行循环外的语句
		break
	}
	//会执行
	fmt.Println("结束之后的语句....")
}
