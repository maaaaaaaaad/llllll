package main

import (
	"fmt"
)

type x interface{
  foo(n uint8) uint8
}

func foo(n uint8) uint8{
  fmt.Println(&n)
  return n
}

func main() {
	var memory uint8
	memory = 7
  fmt.Println(&memory)
  result := foo(memory)
  fmt.Println(&result)
  fmt.Println(&result == &memory)
}
