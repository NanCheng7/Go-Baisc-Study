package main

import "fmt"

func main() {
	var base [30]int

	for i, _ := range base {
		base[i] = i + 1
	}
	//区间
	fmt.Println(base[10:15])
	fmt.Println(base[20:])
	fmt.Println(base[:10])

}
