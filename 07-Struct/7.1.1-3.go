package main

import "fmt"

type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令绑定的变量
	Comment string // 指令的注释
}

func newCommand(name string, varRef *int, comment string) *Command {
	return &Command{
		name,
		varRef,
		comment,
	}
}

var version = 1

func main() {
	cmd := newCommand(
		"version",
		&version,
		"show version",
	)
	version = 3
	fmt.Println(*cmd.Var)
}
