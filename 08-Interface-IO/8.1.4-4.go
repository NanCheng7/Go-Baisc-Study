package main

import "fmt"

type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}
type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println(c.name, ": 喵喵喵")
}

func (c cat) move() {
	fmt.Println(c.name, ": 会动")
}

func main() {
	var x animal
	x = cat{name: "花花"}
	x.move()
	x.say()
}
