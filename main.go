package main

import (
	"fmt"
)

func main() {
	x := 4
	y := 4

	result := square(x, y)
	fmt.Println(result)

}

func square(a int, b int) int {
	return a + b
}
