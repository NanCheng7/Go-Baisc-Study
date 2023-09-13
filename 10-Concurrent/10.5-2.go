package main

import "fmt"

func main() {
	// 创建两个管道
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	go func() {
		// time.Sleep(2 * time。Second)
		intChan <- 1
	}()
	go func() {
		stringChan <- "hello"
	}()
	select {
	case value := <-intChan:
		fmt.Println("int:", value)
	case value := <-stringChan:
		fmt.Println("string:", value)
	}
	fmt.Println("main 结束")
}
