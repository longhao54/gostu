package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func main() {
	two1 := new(TwoInts)
	two1.a, two1.b = 1, 2
	fmt.Println(two1.AddThem())
	fmt.Println(two1.AddToParam(2))
	c := TwoInts{5, 5}
	fmt.Println(c.AddThem())

}
