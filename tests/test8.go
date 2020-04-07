package main

import (
	"fmt"
	"net"
)

// net 库测试

func main() {
	a, b, err := net.SplitHostPort("http://127.0.0.1:2379")
	fmt.Println("1" + a)
	fmt.Println("2" + b)
	fmt.Println(err)
}
