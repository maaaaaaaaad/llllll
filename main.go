package main

import "fmt"

func main() {
	people := []string{"a", "b", "c", "d", "e", "f"}
	for i := 0; i < len(people); i++ {
		for j := i + 1; j < len(people); j++ {
			fmt.Printf("(%s, %s)\n", people[i], people[j])
		}
	}
}
