package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := firstCheckError()
	if err != nil {
		fmt.Println(err)
		exitProcess()
	}
	err = secondCheckError()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")
}
func secondCheckError() interface{} {
	return errors.New("error 2")
}
func firstCheckError() interface{} {
	return errors.New("error 1")
}
func exitProcess() {
	os.Exit(1)
}
