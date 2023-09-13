package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello Goroutine")
}
func main() {
	go hello() // 启动一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)

}
