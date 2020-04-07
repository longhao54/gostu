package main

import (
	"fmt"
	"sync"
)

//  sync.Once 模块测试

var syc_s struct {
	once sync.Once
	a    int
}

func doonce() {
	syc_s.once.Do(func() {
		fmt.Println("do once")
	})
}

func main() {
	doonce()
	doonce()
}
