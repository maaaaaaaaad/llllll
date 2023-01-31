package main

import "fmt"

func main() {
	x := make([]int, 5, 5)
	for i := 0; i < len(x); i++ {
		x[i] = i + 1
	}
	fmt.Println("x:", x)
}
