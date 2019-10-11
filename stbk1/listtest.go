package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var b [5]string
	var c [5]int = [5]int{1, 2, 3}
	var d = [...]int{1, 2, 3} // 长度根据chu
	// hahahahahahadfefasdf
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(c[1])
	fmt.Println(d)
	fmt.Println("asdf")
	for i, v := range a {
		fmt.Println(i, v)
	}
}
