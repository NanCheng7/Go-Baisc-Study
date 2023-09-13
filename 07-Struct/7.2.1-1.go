package main

import "fmt"

// 定义属性结构
type Property struct {
	value int // 属性值
}

// 设置属性值
func (p *Property) SetValue(v int) {
	// 修改p的成员变量
	p.value = v
}

// 取属性值
func (p *Property) Value() int {
	return p.value
}
func main() {
	// 实例化属性
	p := new(Property)
	// 设置值
	p.SetValue(100)
	// 打印值
	fmt.Println(p.Value())
}
