package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {

	w, h := 98, 98
	s := w * h

	ds := 2*(2*98) + 4

	r := (100 * 100) - ds

	if s == r {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}
