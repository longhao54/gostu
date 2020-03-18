package main

import "reflect"

func main() {
	t1 := make(map[reflect.Type]int)
	t1[string] = 2
}
