package main

import "fmt"

func main() {
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	//容量不会影响当前的元素个数，因此a和b取len都是2
	//但是如果我们给a追加一个元素，a的长度就会变成3
	fmt.Println(len(a), len(b))
	a = append(a, 1, 2, 3)
	fmt.Println(len(a))
}
