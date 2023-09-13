package main

import "fmt"

func main() {
	//#region  类型强转
	//coding:
	//fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	//fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	//fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	//fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	//
	//var a int32 = 1047483647
	//
	//fmt.Printf("int32: 0x%x %d \n", a, a)
	//
	//b := int16(a)
	//
	//fmt.Printf("int16: 0x%x %d\n", b, b)
	//
	//var c float32 = math.Pi
	//fmt.Println(int(c))

	/* noting:
	通过上述代码可以知道int的各个位数类型对应的最大最小值
	例如我们在上面的代码中定义了一个a变量，它的类型是32位int
	当我们强转为int8类型，则会发生精度损失
	如果我们将浮点数转换为int类型，则会直接抛弃小数部分
	*/
	//#endregion

	//#region 拼接字符串

	s1 := "localhost:8080"
	bytes := []byte(s1)
	bytes[len(bytes)-1] = '1'
	fmt.Println(string(bytes))

	//#endregion
}
