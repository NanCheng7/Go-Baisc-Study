package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++
		if sum > 100 {
			//break是跳出循环
			break
		}
	}
	fmt.Println(sum)
}
