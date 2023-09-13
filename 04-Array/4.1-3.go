package main

import "fmt"

func main() {
	type ai3 [3]int
	var arr ai3
	arr[0] = 2
	for index, value := range arr {
		fmt.Printf("索引: %d，值: %d \n", index, value)
	}
}
