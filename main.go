package main

import (
	"fmt"
	"math/rand"
	"time"
)

type member struct {
	name string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var key uint8 = uint8(rand.Intn(3) + 1)
	fmt.Println(key)

	var members = []member{
		member{
			name: "mad",
		},
		member{
			name: "integral",
		},
		member{
			name: "mooze",
		},
	}

	fmt.Println(members)
}
