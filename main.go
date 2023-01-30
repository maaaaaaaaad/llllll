package main

import "fmt"

func main() {
	s1 := []uint8{1, 2, 3, 4, 5}
	s2 := make([]uint8, len(s1))

	for i, v := range s1 {
		s2[i] = v
	}

	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2)) // {1,2,3,4,5}, 5, 5

	s2[0] = 255

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}
