package main

type Emp struct {
	IN        int
	name      string
	Address   string
	ManagerID int
}

func mian() {
	var test Emp
	test.IN = 1
	a := &test.name
	*a = "kim"

	type Point struct{ X, Y int }
	type Circle struct {
		Center Point
		Radius int
	}
	var t Circle
	t.Center.X = 2 // 结构体的嵌套

	type Cir struct { // 定义一个匿名结构体
		Point
		radius int
	}
	var tt Cir
	tt.X = 2
}
