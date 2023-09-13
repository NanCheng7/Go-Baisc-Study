package main

import "fmt"

func main() {
	var cat int = 1
	var str string = "南城"
	fmt.Printf("%p %p", &cat, &str)
}
