package main

import "fmt"

func changeA(args [5]uint8) {
	args[0] = 100
}
func changeS(args []uint8) {
	args[0] = 100
}

func main() {
	a := [5]uint8{1, 2, 3, 4, 5}
	s := []uint8{1, 2, 3, 4, 5}

	changeA(a)
	changeS(s)

	fmt.Println(a)
	fmt.Println(s)
}
