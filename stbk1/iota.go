package main

import "fmt"
// iota 用于创建一系列相关值的常量。 iota从0开始取值 逐项+1

type Weekday int
type Flags uint

// iota 的简单使用
const (
	sun Weekday = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
)

const (  // 1, 2, 4
//  _ = 1 << (10 * iota)  1024的幂
    a Flags = 1 << iota // 2的幂
    b
    c
)

func main() {
	fmt.Println(sun, Wed)
    fmt.Println(a,b,c)
}