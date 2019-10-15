package main

import "fmt"
import "math"

type Square2 struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper2 interface {
	Area() float32
}

func (sq *Square2) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper2
	sq1 := new(Square2)
	sq1.side = 5

	areaIntf = sq1

	// 检测 areaIntf的类型是否时 Square2
	if t, ok := areaIntf.(*Square2); ok {
		fmt.Println(1, t)
	}
	if t, ok := areaIntf.(*Circle); ok {
		fmt.Println(1, t)
	}
}
