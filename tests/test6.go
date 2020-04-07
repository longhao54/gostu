package main

import (
	"fmt"
	"os"
)

// func 的执行时间测试

func t6_1(a string) func() {
	b := func() {
		fmt.Print(a)
		fmt.Println(os.Hostname())
		fmt.Println("is os")
	}
	return b
}

func main() {
	a := t6_1("asdfe")
	fmt.Println("1")
	a()
	fmt.Println("2")
}
