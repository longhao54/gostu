package main

import (
	"fmt"
	"time"
)

// struct type channel test

func get_ch() chan struct{} {
	ch := make(chan struct{})
	return ch
}

func run(ch <-chan struct{}) {
	fmt.Println("run start")
	<-ch // 在 close 这个ch以后才会继续进行  struct类型的channel不占用任何内存
	fmt.Println("run stop")

}

func test1(ch <-chan struct{}) {

}

func main() {
	ch := get_ch()
	go run(ch)
	time.Sleep(10 * time.Second)
	close(ch)
	fmt.Println("closed")
	fmt.Println("cl")
}
