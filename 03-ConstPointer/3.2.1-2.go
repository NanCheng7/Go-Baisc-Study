package main

import "fmt"

func main() {
	// 指针与变量
	var room int = 10 // room房间 里面放的 变量10
	var ptr = &room   // 门牌号px  指针  0xc00000a0a8

	fmt.Printf("%p\n", &room) // 变量的内存地址 0xc00000a0a8

	fmt.Printf("%T, %p\n", ptr, ptr) // *int, 0xc00000a0a8

	fmt.Println("指针地址", ptr)      // 0xc00000a0a8
	fmt.Println("指针地址代表的值", *ptr) // 10
}
