package main

import (
	"fmt"
)

func main() {
	var arr = [3]int{1, 2, 3}
	var arr1 = [3]int{1, 2}
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3}
	printArr(arr)
	fmt.Println("---------------")
	printArr(arr1)
	fmt.Println("---------------")
	printArr(arr2)
	fmt.Println("---------------")
	printArr(arr3)

}
func printArr(array [3]int) {
	for index, var1 := range array {
		fmt.Printf("%d => %d\n", index, var1)
	}
}
