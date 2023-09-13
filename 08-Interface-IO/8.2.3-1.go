package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wr() {
	//参数2：打开模式
	//参数3：权限控制
	//w 写 r 读 x 执行  w：2  r：4 x：1
	//特殊权限为，拥有者位，同组用户位，其余用户为
	file, err := os.OpenFile(".\\MSZL\\Study\\08-Interface-IO\\8.2.3-1.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Open Error : ", err)
		return
	}
	defer file.Close()
	// 获取wirter对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	//刷新缓冲区，强制写出
	writer.Flush()
}

func re() {
	file, err := os.Open(".\\MSZL\\Study\\08-Interface-IO\\8.2.3-1.txt")
	if err != nil {
		fmt.Println("Open Error : ", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("Read Over")
			break
		}
		if err != nil {
			fmt.Println("Read Error : ", err)
			return
		}
		fmt.Println(string(line))
	}
}
func main() {
	re()
	wr()
}
