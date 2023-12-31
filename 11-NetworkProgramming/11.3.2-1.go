package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http://127.0.0.1:8000/go
	//单独写回调函数

	http.HandleFunc("/go", myHandler)
	//addr：监听的地址
	//handler：回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// handler 函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	//request method： GET POST DELETE PUT UPDATE
	fmt.Println("method", r.Method)
	// /go
	fmt.Println("url", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	//回复
	w.Write([]byte("Hello,NanCheng"))
}
