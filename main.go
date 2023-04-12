package main

import "fmt"

const head = 4
const diff = 9

func sequence(n uint) uint {
	if n == 1 {
		return head
	} else {
		return sequence(n-1) + diff
	}
}

func main() {
	result := sequence(5)
	fmt.Println(result)
}
