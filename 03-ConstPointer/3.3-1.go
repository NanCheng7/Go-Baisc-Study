package main

import "fmt"

var global *int

func f() {
	var x int
	x = 1
	global = &x
}
func g() {
	y := new(int)
	*y = 1
	fmt.Println("&y:", y, "\ny:", *y)
}

func main() {
	f()
	fmt.Println("&x:", global, "\nx:", *global)
	g()
}
