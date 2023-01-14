package main

import (
	"fmt"
)

func cone(r int, h int) int {
	return (r * r * h) / 3
}

func main() {

	r := 98
	h := 98

	s := cone(r, h)

	fmt.Println(s)

}
