package main

import "fmt"

func main() {
	a := [5]uint8{1, 2, 3, 4, 5} // len 5, cap 5
	fmt.Printf("a: %v, len: %d, cap: %d\n", a, len(a), cap(a))
	fmt.Printf("a memory: %p\n", &a)
	s := a[1:2] // len 1, cap 4
  fmt.Printf("s memory: %p\n", &s)

	s[0] = 100 // s = [100], len 1, cap 4

	s = append(s, 200, 250, 255) // s = [100, 200, 250, 255], len 4, cap 4
	s = append(s, 180)           // s = [100, 200, 250, 255, 180], len 5, cap 8
	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s))
	fmt.Println("a:", a)
}
