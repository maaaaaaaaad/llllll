package main

import (
	"fmt"
)

func rpointer(a *int, b *int, f func(*int, *int) (int, int)) (int, int) {
	return f(a, b)
}

func lounge(a *int, b *int) (r1 int, r2 int) {
  r1 = *a * 50
  r2 = *b * 20
  return r1, r2
}

func main() {
	p, q := 5, 6
  w, z := rpointer(&p, &q, lounge)
  fmt.Println(w, z)
}
