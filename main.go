package main

import (
	"fmt"
)

func rpointer(a *int, b *int) (r1 int, r2 int) {
	r1 = *a * 100
	r2 = *b * 100
	return r1, r2
}

func main() {
	p, q := 5, 6
	result1, result2 := rpointer(&p, &q)
	fmt.Println(result1, result2)
}
