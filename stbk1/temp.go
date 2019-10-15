package main

import "person"
import "fmt"

func main() {
	p := new(person.Person)
	p.SetFirstName("tom")
	fmt.Println(p.FirstName())
}
