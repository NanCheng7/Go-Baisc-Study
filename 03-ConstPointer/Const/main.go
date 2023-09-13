package main

import "fmt"

func main() {

	//#region 常量定义

	//coding:
	//const (
	//
	//	a = 1
	//	b
	//	c
	//	d
	//	e = 2
	//	f
	//	g
	//	h = "hello"
	//)
	//fmt.Println(a, b, c, d, e, f, g, h)

	/* noting:
	我们在进行常量的批量定义时，如果后续定义变量为赋值，则默认和前面最后一个变量的赋值相同
	*/

	//#endregion

	//#region iota 常量生成器

	//coding:
	const (
		Sunday = iota //1
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday //6
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	/* noting:
	第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加1
	*/

	//#endregion

}
