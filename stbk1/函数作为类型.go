package main

import "fmt"

func isOdd(v int) bool {
	if v%2 == 0 {
		return false
	}
	return true
}

func isEven(v int) bool {
	if v%2 == 0 {
		return true
	}
	return false
}

type boolFunc func(int) bool // 声明一个函数类型 当作一个参数使用
// 可以作为其他函数中的 通用接口函数变量使用

func filter(slice []int, f boolFunc) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd: ", odd)
	even := filter(slice, isEven)
	fmt.Println("even: ", even)
}
