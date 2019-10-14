// like dic
// map 中 k需要相同的数据类型  v 也需要相同的数据类型

package main

import "fmt"

func main() {
	var age map[string]int
	test := map[string]int{
		"alice": 31,
		"kim":   20,
	}
	fmt.Println(age, test)
	delete(test, "test") // 从map中删除 k

	for name, age := range test {
		fmt.Printf("%s \t %d \n", name, age)
	}

	var prereqs = map[string][]string{
		"a": {"b", "c"},
	}
	fmt.Println(prereqs)
}
