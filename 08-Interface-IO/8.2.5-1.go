package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') // 注意是字符
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
	}
}
func main() {
	flag.Parse() //解析命令行参数
	if flag.NArg() == 0 {
		//如果没有默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	//依次读取每个指定文件的内容并且打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "Reading from %s failed, err:^v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}
