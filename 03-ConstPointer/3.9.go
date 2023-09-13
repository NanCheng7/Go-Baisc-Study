package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var level = 1
var ex = 0

func main() {
	fmt.Println("请输入你的角色名字：")
	//#region 捕获标准输入，并转换为字符串
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	//#endregion
	//删除最后的\n
	name := input[:len(input)-2]
	fmt.Printf("角色创建成功,%s,欢迎你来到异世界,目前角色等级为%d级", name, level)
	s := `你遇到了一个怪物，请选择战斗还是逃跑?
	1.战斗
	2.逃跑`
	fmt.Println(s)
	var e int
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		selector := input[:len(input)-2]
		switch selector {
		case "1":
			e = rand.Intn(20)
			ex += e
			fmt.Println("杀死了怪物，获得了", e, "经验")
			computeLevel()
			fmt.Println("您现在的等级为", level, "级\n您现在的经验为", ex)
		case "2":
			fmt.Println("你选择了逃跑")
			fmt.Println(s)
		case "exit":
			fmt.Println("你退出了游戏")
			os.Exit(1)
		default:
			fmt.Println("无法识别宿主命令，请重新输入")
		}
	}
}
func computeLevel() {
	if ex >= 20 {
		ex -= 20
		level++
	}
}
