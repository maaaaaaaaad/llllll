package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	n := m
	n["one"] = 5

	p := make([]int, 5, 5)
	q := p
	q[0] = 10
	fmt.Println(p, q)

	x := [3]int{1, 2, 3}
	y := x
	y[0] = 10
	fmt.Println(x, y)
}
