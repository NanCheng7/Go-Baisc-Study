package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8
	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24
	var m map[int]bool
	fmt.Println(unsafe.Sizeof(m)) // 8
	var c chan string
	fmt.Println(unsafe.Sizeof(c)) // 8
	var f func()
	fmt.Println(unsafe.Sizeof(f)) // 8
	var i interface{}
	fmt.Println(unsafe.Sizeof(i)) // 16
}
