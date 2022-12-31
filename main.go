package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	mad := person{}
	mad.name = "mad"
	mad.age = 31
	fmt.Println(mad)
}
