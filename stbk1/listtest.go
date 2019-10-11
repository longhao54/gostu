package main

import "fmt"

// 不同长度的数组 无法进行比较.  会报错
func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var b [5]string
	var c [5]int = [5]int{1, 2, 3}
	var d = [...]int{1, 2, 3} // 长度根据传入的数据决定
	var e = [...]int{99: 2}   // 长度为100的数组. 前99项值为0 最后一个值为2
	// hahahahahahadfefasdf
	fmt.Println(len(a))
	fmt.Println(len(b), e)
	fmt.Println(len(c), len(d))

	for i, v := range a {
		fmt.Println(i, v)
	}

	type currency int
	const (
		USD currency = iota
		EUR
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "aa", RMB: "#"} // 通过在上面用了iota 自增常量 . 相当于给数据每个下标都取了一个别名
	fmt.Println(RMB, symbol[RMB])

}
