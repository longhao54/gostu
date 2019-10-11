package main

import "fmt"
import "strconv"

// int 转 string 的两种方式

const changliang = 1

const (
	changliang1 = 1
	changliang2 = 2
)

func main() {
	x := 123
	a := fmt.Sprintf("%d", x)

	fmt.Println(a, strconv.Itoa(x))
}