package main

import (
	"fmt"
)

type bundle interface {
	checkInt(n interface{}) bool
}

func checkInt(n interface{}) bool {
	switch n.(type) {
	case int:
		return true
	default:
		return false
	}
}

func main() {
	n, m := 255, "hi"
	r1 := checkInt(n)
	r2 := checkInt(m)
  fmt.Print("r1:", r1)
  fmt.Print("r1:", r1)
}
