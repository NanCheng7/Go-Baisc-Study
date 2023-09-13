package main

import (
	"fmt"
	"strconv"
)

func main() {
	//str <=> float
	str := "3.1415926"
	f, _ := strconv.ParseFloat(str, 32)
	fmt.Printf("%T,%f \n", f, f)
	//float <=> str
	floatValue := 3.1415926
	/*
		4个参数
		- 要转换的浮点数
		- 格式标记 (b,e,E,f,g,G)
			- 'b' 二进制指数
			- 'e' 十进制指数
			- 'E'十进制指数
		 	- 'f' 没有指数
			- 'g' 大指数
			- 'G' 大指数
		- 精度 如果格式标记是 e、E、f,则精度表示小数点后的数字位数，否则表示整数+小数部分的数字位数
		- 指定浮点类型 (32: float32;64:float64)
	*/
	formatFloat := strconv.FormatFloat(floatValue, 'f', 4, 64)
	fmt.Printf("%T,%s", formatFloat, formatFloat)

}
