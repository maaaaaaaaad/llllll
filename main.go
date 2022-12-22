package main

import (
	"fmt"
)

func main() {
	const (
		A = iota + 0.75 + 0.25
		B
		C
		D
	)
	fmt.Println(A + B + C + D)
}
