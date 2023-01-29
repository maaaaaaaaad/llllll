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

	p := []uint8{1, 2, 3}

	q := append(p, 4, 5)
	fmt.Println("q", q)

  p = append(p, 100)
  fmt.Println("p", p)
}
