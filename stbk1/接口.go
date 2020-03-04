package main

import "fmt"

type Shaper interface {
	Area() float64
}

type Square struct {
	side float64
}

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

func main() {
	//sq1 := new(Square)
	//sq1.side = 5

	var sql Shaper = &Square{5}
	fmt.Println(sql.Area())

	//areaIntf := sq1
	//fmt.Println(areaIntf.Area())
}
