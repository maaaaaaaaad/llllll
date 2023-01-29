package main

import "fmt"

func main() {
	a := [5]uint8{1, 2, 3, 4, 5}
	s := a[1:2]
	fmt.Println(s)

	s[0] = 100
	fmt.Println("a", len(a), cap(a), a)
	fmt.Println("s:", len(s), cap(s), s)

	s = append(s, 200, 250, 255)
	fmt.Println("s:", len(s), cap(s), s)
}
