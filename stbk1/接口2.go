package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float64
	count      float64
}

type car struct {
	make   string
	modell string
	price  float64
}

type valuable interface {
	getValue() float64
}

func (s stockPosition) getValue() float64 {
	return s.sharePrice * s.count
}

func (c car) getValue() float64 {
	return c.price
}

func showValue(asset valuable) {
	fmt.Println(asset.getValue())
}

func main() {
	var o valuable = stockPosition{"a", 5, 20}
	showValue(o)
	o = car{"b", "c", 5}
	showValue(o)

}
