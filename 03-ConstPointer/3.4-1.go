/*
运行包
*/
package main

import "fmt"

type NewInt int
type IntAlias = int

func main() {
	//a => NewInt
	var a NewInt
	//查看a 的类型名
	fmt.Printf("a => type: %T\n", a)
	//b => IntAlias
	var b IntAlias
	//查看b 的类型名
	fmt.Printf("b => type: %T\n", b)
}
