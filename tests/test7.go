package main

import "fmt"

// return 值的测试

type c struct {
	a int
}

func retest() (rc c) {
	rc = c{1}
	return
}

func main() {
	c := retest()
	fmt.Println(c)
}
