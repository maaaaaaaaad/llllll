package main

import (
	"fmt"
)

type mock struct {
	first  int
	second int
}

type functions interface {
	foo() int
	bar() int
}

func (receiver mock) foo() int {
	return receiver.first
}
func (receiver mock) bar() int {
	return receiver.second
}

func call(arg functions) (int, int) {
	a := arg.foo()
	b := arg.bar()
	return a, b
}

func main() {
	n := mock{
		first:  1,
		second: 2,
	}

	r1, r2 := call(n)
	fmt.Println(r1, r2)
}
