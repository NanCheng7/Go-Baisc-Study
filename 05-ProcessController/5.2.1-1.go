package main

import "fmt"

func length(s string) int {
	println("call length.")
	return len(s)
}
func main() {
	s := "acbd"
	//多次调用length函数
	for i := 0; i < length(s); i++ {
		print(i)
		fmt.Printf("%c\n", s[i])
	}

}
