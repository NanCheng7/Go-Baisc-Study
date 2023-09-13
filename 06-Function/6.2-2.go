package main

import "fmt"

func main() {
	func(data int) {
		fmt.Println("hello", data)
	}(100) //(100)，表示对匿名函数进行调用，传递参数为 100。
}
