package main

import "fmt"

func main() {
	s1 := []uint8{1, 2, 3, 4, 5}
	p := make([]uint8, len(s1), cap(s1))

	s2 := copy(p, s1)
	fmt.Println(s2)
	fmt.Println("p:", p)
}
