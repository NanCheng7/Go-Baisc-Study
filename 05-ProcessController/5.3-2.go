package main

import "fmt"

func main() {
	str := "南城QJ"
	//因为一个字符串是 Unicode 编码的字符（或称之为 rune ）集合
	//char 实际类型是 rune 类型
	for pos, char := range str {
		fmt.Printf("%d,%c\n", pos, rune(char))
	}
}
