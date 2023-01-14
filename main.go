package main

import (
	"fmt"
	"math/rand"
	"time"
)

type person struct {
	name string
}

func main() {

	m := []person{
		{
			name: "mad",
		},
		{
			name: "intel",
		},
		{
			name: "gogo",
		},
		{
			name: "gral",
		},
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(3 + 1)
	fmt.Println(m[random])
}
