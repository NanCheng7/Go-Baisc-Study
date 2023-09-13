package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	//coding:
	s1 := "hello,南城你好呀"
	s2 := "在学Go语言"

	target := "好"
	index := strings.Index(s1, target)

	s01, s02 := s1[:index], s1[index+len(target):]

	var sb bytes.Buffer

	sb.WriteString(s01)
	sb.WriteString(s2)
	sb.WriteString(s02)
	fmt.Println(sb.String())

	/* noting:
	通过上面的代码，我们可以很清晰地看到，字符串替换实际上就是利用了切片，将字符串切割为几个部分后
	替换掉不需要的部分，再将需要的部分插入，最后进行拼接
	*/
}
