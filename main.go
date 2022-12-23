package main

import (
	"fmt"
	"main/lib"
)

func main() {
	member := []string{"mad", "woong", "integral"}
	result := lib.Random(member)
	fmt.Println(result)
}
