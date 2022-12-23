package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var key = []string{"mad", "woong", "integral"}
	var index = rand.Intn(3)
	result := key[index]
	fmt.Println(result)
}
