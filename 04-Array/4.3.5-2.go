package main

import "fmt"

func main() {
	mp1 := make(map[int][]int)
	mp2 := make(map[int]*[]int)

	mp1 = map[int][]int{1: []int{1, 2, 3}, 2: {4, 5, 6}}
	mp2 = map[int]*[]int{1: {mp1[1][1]}}

	fmt.Println(mp1)
	fmt.Println(mp2)
}
