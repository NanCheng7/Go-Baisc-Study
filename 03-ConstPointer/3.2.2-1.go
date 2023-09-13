package main

import "fmt"

func main() {
	// 利用指针修改值
	var num = 10
	modifyFromPoint(num)
	fmt.Println("未使用指针，方法外", num)

	var num2 = 22
	newModifyFromPoint(&num2) // 传入指针
	fmt.Println("使用指针 方法外", num2)
}

func modifyFromPoint(num int) {
	// 未使用指针
	num = 10000
	fmt.Println("未使用指针，方法内:", num)
}

func newModifyFromPoint(ptr *int) {
	// 使用指针
	*ptr = 1000 // 修改指针地址指向的值
	fmt.Println("使用指针，方法内:", *ptr)
}
