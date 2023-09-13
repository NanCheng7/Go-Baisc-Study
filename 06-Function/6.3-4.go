package main

import "fmt"

func main() {
	var whatever = [5]int{1, 2, 3, 4, 5}
	for i, _ := range whatever {
		//函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.
		defer func() { fmt.Println(i) }()
	}
}
