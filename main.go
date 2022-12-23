package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	sum2, j := 0, 0
	for j <= 100 {
		sum2 += j
		j++
	}
	fmt.Println("sum2:", sum2)
}
