package main

import "fmt"

func test() {
	defer func() {
		// defer panic 会打印
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

func main() {
	test()
}
