package main

import (
	"fmt"
)

func main() {
	//invalid operation: nil == nil (operator == not defined on nil)
	fmt.Println(nil == nil)
}
