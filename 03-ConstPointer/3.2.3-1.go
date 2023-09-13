package main

import "fmt"

func main() {
	str := new(string)
	*str = "南城"
	fmt.Println(*str)
}
