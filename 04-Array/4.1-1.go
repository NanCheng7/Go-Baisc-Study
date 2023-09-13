package main

import "fmt"

func main() {
	var arr [3]int
	for index, value := range arr {
		fmt.Printf("索引:%d，值:%d \n", index, value)
	}
}
