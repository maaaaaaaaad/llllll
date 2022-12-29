package main

import (
	"fmt"
	"main/lib"
)

func chance(a int, b int, f func(int, int) (r1 int, r2 int)) (x int, y int) {
	x, y := f(a,b)
}

func change(a *int, b *int) (r1 int, r2 int) (r1 int, r2 int) {
  r1 = *a * 100
  r2 = *b * 100
  return r1, r2
}

func main() {
	p, q := 10
  }
