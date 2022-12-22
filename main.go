package main

import (
	"fmt"
)

func main() {
	names := []string{"a", "b", "c"}

	for index, name := range names {
		fmt.Println(index, name)
	}
}
