package main

import "fmt"

func main() {

	scene := make(map[string]int)
	// 准备map数据
	scene["cat"] = 66
	scene["dog"] = 4
	scene["pig"] = 960
	delete(scene, "dog")
	for k, v := range scene {
		fmt.Println(k, v)
	}
}
