// 更准确的是应该是当成一个 长度不固定的数组 可以用 pyhton当中的切片操作
// slice 无法用 == 做比较 判断里面的函数是否相同
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5} // this is slice
	a := [...]int{1, 3, 4, 5} // this is list
	b := [3]int{1, 2, 3}      // this is list

	fmt.Println(a, b, s)

	var runes []rune

	for _, r := range "hello world" {
		runes = append(runes, r)
	}
	fmt.Println(runes)
	fmt.Printf("%q\n", runes)
	fmt.Println(cap(runes))
}
