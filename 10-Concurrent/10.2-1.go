package main

import "fmt"

func hello() {
	fmt.Println("hello Goroutine")
}

func main() {
	hello()
	fmt.Println("main hello Goroutine")
}
