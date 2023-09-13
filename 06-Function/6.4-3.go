package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	log.Printf("开始时间为：%v", start)
	defer func() {
		log.Printf("开始调用defer")
		log.Printf("时间差：%v", time.Since(start))
		log.Printf("结束调用defer")
	}()
	time.Sleep(3 * time.Second)

	log.Printf("函数结束")
}
