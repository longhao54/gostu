package main

import "fmt"

// type xxxx func()   测试
// 当作一个 参数来使用就行了 或者看 test3的那种用法也行

type st func(a string, b string) string

func stused(s1 string, s2 string) string {
	fmt.Println(s1)
	fmt.Println(s2)
	return s1 + s2
}

func t1(st1 st) {
	fmt.Println(st1("asdf", "asdf"))
}

func main() {
	t1(stused)
}
