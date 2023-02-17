package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	members := []string{
		"mad",
		"integral",
		"woong",
	}

	random := uint8(rand.Intn(len(members)))
	fmt.Println(members[random])
}
