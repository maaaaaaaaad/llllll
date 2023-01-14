package main

import (
	"fmt"
)

func main() {
	var memory int
	memory = 4
	fmt.Printf("memory address %p, value %d", &memory, memory)
}
