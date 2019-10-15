package main

import "fmt"

type Shaper1 interface {
	Area() float64
}

type Square1 struct {
	side float64
}

func (sq *Square1) Area() float64 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3}
	q := &Square1{5}
	shapes := []Shaper1{r, q}
	for n, _ := range shapes {
		fmt.Println(n)
		fmt.Println(shapes[n])
		fmt.Println(shapes[n].Area())
	}
}
