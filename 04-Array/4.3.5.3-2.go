package main

import (
	"fmt"
	"sync"
)

func main() {
	//sync.Map 不能使用make创建
	var scene sync.Map
	//将键值对保存到sync.Map
	//sync.Map 将键和值以 interface{} 类型进行保存
	scene.Store("NanCheng", 1001)
	scene.Store("NanCheng1", 1001)
	scene.Store("NanCheng2", 1001)
	//从 sync.Map中根据键取值
	fmt.Println(scene.Load("NanCheng"))
	//根据键删除对应的键值对
	scene.Delete("NanCheng")
	//遍历所有sync.Map中的键值对
	//遍历需要提供一个匿名函数，参数为k、v，类型为 interface{}，
	//每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}
