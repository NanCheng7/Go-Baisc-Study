package main

import (
	"fmt"
)

func recv(c chan int) {
	ret := <-c
	fmt.Println("recieve access", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch) //启动gorouine从通道接收
	ch <- 10
	//runtime.Gosched()
	fmt.Println("send access")
}
