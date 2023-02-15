package main

import (
	"fmt"
)

func checkSlice(args []int) {
	for i := 0; i < len(args); i++ {
		args[i] += i + 10
	}
}

func main() {
	var a = make([]int, 5, 5)
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	var b = make([]int, len(a), cap(a))
	copy(b, a)
	fmt.Println("a", a)
	fmt.Println("b", b)

	b[0] = 100

	fmt.Println("after a", a)
	fmt.Println("after b", b)

	checkSlice(b)
	fmt.Println("after b", b)
}
