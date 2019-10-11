package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string{
	var buf bytes.Buffer  // bytes.buffer 无须初始化。 起始为空
	buf.WriteByte('[')
	fmt.Println(values)

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	fmt.Println(&buf)
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1,2,3}))
}