package main

import "fmt"

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice := numbers4[4:6]
	//这打印出来长度为2
	fmt.Printf("myslice为 %d, 其长度为: %d\n", mySlice, len(mySlice))
	fmt.Println(&mySlice)
	fmt.Println(cap(mySlice))
	//mySlice = mySlice[:cap(mySlice)]
	mySlice = mySlice[:3]
	fmt.Println(&numbers4)
	fmt.Println(&mySlice)
	//为什么 mySlice 的长度为2，却能访问到第四个元素
	fmt.Printf("myslice的第四个元素为: %d", mySlice[3])
}
