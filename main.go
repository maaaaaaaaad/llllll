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

func febSequnce(n uint) uint {
	if n == 1 || n == 2 {
		return 1
	} else {
		return sequence(n-1) + sequence(n-2)
	}
}

func main() {
	result := sequence(5)
	fmt.Println(result)

	feResult := febSequnce(5)
	fmt.Println(feResult)
}
