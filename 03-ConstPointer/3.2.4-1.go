package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "fast模式能让程序运行的更快")

func main() {
	//解析命令行参数
	flag.Parse()
	fmt.Println(*mode)
}
