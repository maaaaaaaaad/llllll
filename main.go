package main

import "fmt"

func main() {
	s1 := []uint8{1, 2, 3, 4, 5}
	p := make([]uint8, len(s1), cap(s1))

	s2 := copy(p, s1)
	fmt.Println(s2)
	fmt.Println("p:", p)

	s3 := make([]uint8, 5, 5)
	for i := 0; i < len(s3); i++ {
		s3[i] = uint8(i + 1)
	}
	fmt.Println("s3:", s3)

	s4 := make([]uint8, len(s3), cap(s3))
	copy(s4, s3)
	s4[0] = 255
	fmt.Println("s4:", s4)
}
