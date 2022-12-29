package main

import (
	"fmt"
)

func rpointer(a *int, b *int, f func(*int, *int) (int, int)) {
	x, y := f(a, b)
  fmt.Println(x, y)
}

func lounge(a *int, b *int) (r1 int, r2 int) {
  r1 = *a * 100
  r2 = *b * 100
  return r1, r2
}

func main() {
	p, q := 5, 6
  rpointer(&p, &q, lounge)
}
