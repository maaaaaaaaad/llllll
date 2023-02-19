package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)
	m[0] = Product{Name: "a", Price: 100}
	m[1] = Product{Name: "b", Price: 200}
	m[2] = Product{Name: "c", Price: 300}

	for i, v := range m{
    fmt.Println(i, v)
  }
}
