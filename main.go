package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	human := [4]string{"one", "two", "three", "four"}
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(4)
	r := human[random]
	fmt.Println(r)
}
