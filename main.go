package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	m := []int{1, 2, 3, 4}
	for _, value := range m {
		fmt.Println(value)
	}
}
