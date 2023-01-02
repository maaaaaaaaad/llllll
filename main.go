package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (person) greeting() string {
	return "hello world!"
}

func main() {
	mad := person{
		name: "mad",
		age:  31,
	}
	welcome := mad.greeting()
	fmt.Println(welcome)
}
