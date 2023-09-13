package main

import (
	"fmt"
)

func main() {
	//error :use of untyped nil
	fmt.Printf("%T", nil)
	print(nil)
}
