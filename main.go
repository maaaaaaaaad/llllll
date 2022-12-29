package main

import (
	"fmt"
)

func change[T int | int8 | int16 | int32 | int64](a *T, b *T) (r1 T, r2 T) {
	return *a * 10, *b * 5
}

func main() {
	p, q := 5, 6
	result1, result2 := change(&p, &q)
	fmt.Println(result1, result2)
}
