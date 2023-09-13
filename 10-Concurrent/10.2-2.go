package main

import "fmt"

func hello() {
	fmt.Println("hello Goroutine")
}
func main() {
	go hello()
	fmt.Println("main hello Routine")
}
