package main

// 测试 go func

import (
	"fmt"
	"time"
)

func test() {
	for {
		fmt.Println("new go func")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("haha")
	go func() {
		for {
			fmt.Println("in func")
			time.Sleep(1 * time.Second)
		}
	}()
	go test()
	time.Sleep(5 * time.Second)
	fmt.Println("out gofunc")
}
