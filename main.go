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

func foo() int {}
func bar() int {}

func call(arg functions) (int, int) {
	a := foo()
	b := bar()
	return a, b
}

func main() {
  
}
