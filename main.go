package main

import (
	"fmt"
)

func main() {
	var a = make([]int, 5, 5)
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	fmt.Println(a)
}
