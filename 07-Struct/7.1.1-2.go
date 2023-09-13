package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	//可以使用
	var p = Point{
		X: 1,
		Y: 2,
	}
	//p = Point{
	//	1,
	//	2,
	//}
	fmt.Printf("%v,x=%d,y=%d", p, p.X, p.Y)
}
