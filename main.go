package main

import (
	"fmt"
)

func main() {
	name := []string{"mad", "woong", "integral"}
	for index, name := range name {
		fmt.Println(index, name)
	}
}
