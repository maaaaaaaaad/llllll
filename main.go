package main

import "fmt"

func main() {
	s1 := []uint8{1, 2, 3}
	s2 := append(s1, 4, 5)

	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("s2:", s2, len(s2), cap(s2))

	s1[0] = 255

	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("s2:", s2, len(s2), cap(s2))

	s1 = append(s1, 255)

	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("s2:", s2, len(s2), cap(s2))
}
