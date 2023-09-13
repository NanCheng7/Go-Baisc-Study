package main

import "fmt"

func main() {
	var whatever = [5]int{1, 2, 3, 4, 5}

	for i := range whatever {
		defer fmt.Println(i)
	}
	defer fmt.Println(666)
	fmt.Println(156)
}
