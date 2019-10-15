package main

import "fmt"
import "math"

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// 用上一层的方法
func (n *NamePoint) Abs2() float64 {
	return n.Point.Abs() * 100
}

type NamePoint struct {
	Point
	name string
}

func main() {
	n := &NamePoint{Point{3, 4}, "kim"}
	//n := NamePoint{Point{3,4}, "kim"}
	fmt.Println(n.Abs())
	fmt.Println(n.Abs2())
}
