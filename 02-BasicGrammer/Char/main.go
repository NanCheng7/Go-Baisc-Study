package main

import (
	"fmt"
	"strings"
)

func main() {

	//#region 测试1 字符串和byte数组

	//var myStr01 string = "hello"
	//var myStr02 [5]byte = [5]byte{104,101,108,108,111}
	//fmt.Printf("myStr01:%s\n",myStr01)
	//fmt.Printf("myStr02:%s\n",myStr02)
	//#endregion

	//#region 测试2 中文字符和 ACII字符(个数)

	//一个字母占一个字节，一个中文占3个字节

	//var myStr01 string = "hello,码神之路"
	//fmt.Printf("len=%d \n", len(myStr01))
	//fmt.Println(myStr01[6])

	/*
		我们在使用下标获取中文字符时获取的只是中文的某个字节而不是整个中文
		我们使用len()方法去获取包含中文的字符串时，对字符个数的数目统计时
		中文字符算3个
		但是我们可以通过utf8.RuneCountInString()函数去统计具体的字符
	*/
	//s := "南城很帅哦"
	//sLength := utf8.RuneCountInString(s)
	//fmt.Println(sLength)

	//#endregion

	//#region 测试3 字符串拼接

	//这种拼接方式性能不高

	//s := "第一部分"+"第"+
	//	"二部分"
	//fmt.Println(s)

	//另一种拼接方式 效率高 类似于Java的StringBuilder

	//s1 := "hello"
	//s2 := "NanCheng"
	//var stringBuilder bytes.Buffer
	//stringBuilder.WriteString(s1)
	//stringBuilder.WriteString(s2)
	//fmt.Println(stringBuilder.String())

	/*
		对于rune字符串,我们不能直接采用下标方式去获取单个字符，需要将字符串强转为Rune数组之后再去获取下标
	*/
	//#endregion

	//#region 测试4 如何获取中文字符串单个字符

	//coding:

	//s1 := "hello"
	//u := s1[0]
	//fmt.Printf("%c", u)
	//
	//fmt.Println("\n-------------------------")
	//
	//s2 := "hello 南城"
	//fmt.Printf("%s", string([]rune(s2)[6])) //字符变成了字符串
	//fmt.Printf("%c", []rune(s2)[6]) //字符
	//fmt.Println([]rune(s2)[6]) //字符在unicode中的编码
	//fmt.Println(string([]rune(s2)[6])) //字符在unicode中的编码

	/* noting:
	直接的索引对于rune类型无效，我们可以使用String方法进行转化
	准确来说：
		通过[]rune(String s)[index] 将字符串转换为rune数组
		再将rune中的单个字符串转换为String进行输出
		注意： rune是byte数组，如果不转换为字符串打印则会直接输出该字符在unicode编码集中的值
	*/

	//#endregion

	//#region 测试5 字符串遍历

	//coding:
	//var str1 string = "hello"
	//var str2 string = "hello 南城"
	//
	//fmt.Println("str1 normal for cycle")
	//for i := 0; i < len(str1); i++ {
	//	fmt.Printf("ancii: %c %d\n", str1[i], str1[i])
	//}
	//
	//fmt.Println("str1 range for cycle")
	//for _, s := range str1 {
	//	fmt.Printf("unicode: %c %d\n", s, s)
	//}
	//
	//fmt.Println("str2 rang for cycle")
	//for _, s := range str2 {
	//	fmt.Printf("unicode: %c %d\n", s, s)
	//}

	/* noting:
	ANCII字符集可以使用 for range或者for循环遍历
	unicode 字符集只能够使用for range 遍历
	*/

	//#endregion

	//#region 测试6 字符串的格式化

	//coding:
	//str1 := "你好，我是南城"
	//result := fmt.Sprintf("字符串是：%s", str1)
	//fmt.Println(result)

	/* noting:
	我们使用printf()可以进行字符串的标准化输出，通过使用占位符来达到字符串数值替换的目的
	当然如果我们想要通过format来进行字符串拼接，也是可以实现的
	使用fmt包下的Sprintf(format,var)可以实现返回字符串的拼接值
	*/
	//#endregion

	//#region 测试7 字符串查找
	main := "南城是个大帅哥,南城直接无敌起飞"

	target1 := "南城"
	index1 := strings.Index(main, target1)
	fmt.Println("the word of", target1, "'s first index is :", index1)

	target2 := "丑逼"
	index2 := strings.Index(main, target2)
	fmt.Println("the word of ", target2, "'s first index is :", index2)

	target3 := "南城"
	index3 := strings.Index(main, target3)
	fmt.Println("the word of ", target3, "'s first index is :", index3)

	target4 := "南城"
	comma := strings.Index(main, ",")
	index4 := strings.Index(main[comma:], target4)
	fmt.Println("in", main[comma:], " the word of ", target3, "'s first index is :", index4)
	fmt.Println(comma, index4, main[5+index4:])
	fmt.Println(main[:comma])

	/* noting:
	在本次测试中，我们学会了字符串中的查找和选取
		首先是strings.index(str1,str2)方法,其作用是在str1中查找str2首次出现的位置，起始位置为0
		其次则是数组的切片,我们使用 main[index:] 则其内容是下标从index其到数组结束的内容(包括index)
		如果想要index之前的内容,则使用main[:index] 则是从数组起始位到下标index(不包括index)
	attention:
		中文字符占3个字节!!!!
	*/

	//#endregion

}
