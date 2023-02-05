package main

import (
	"fmt"
)

func checkType(n interface{}) {
	switch t := n.(type) {
	case uint8:
		fmt.Println(t)
	default:
		fmt.Println("n is not uint8")
	}
}

func main() {
	var m uint8 = 255
	checkType(m)
}
