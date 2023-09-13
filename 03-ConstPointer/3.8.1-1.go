package main

import (
	"fmt"
	"strconv"
)

func main() {
	newStr := "1"
	intValue, _ := strconv.Atoi(newStr)
	fmt.Printf("%T,%d\n", intValue, intValue)

	intValue2 := 1
	strValue := strconv.Itoa(intValue2)
	fmt.Printf("%T,%s\n", strValue, strValue)
}
