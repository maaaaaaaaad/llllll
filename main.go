package main

import (
	"fmt"
)

type per struct {
	name string
	age  uint8
}

type met interface {
	upgrade(n uint8)
}

func (receiver *per) met(n uint8) {
	receiver.age += n
}

func main() {
	m := per{
		name: "mad",
		age:  20,
	}

	fmt.Println(m)

	m.met(2)
	fmt.Println(m)
}
