package main

import (
	"fmt"
)

type checkfields struct {
	m int
	n string
	l float64
}

func check(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("t is %d\n", t)
	case string:
		fmt.Printf("t is %s\n", t)
	case float64:
		fmt.Printf("t is %f\n", t)
	default:
		fmt.Println("default:", t)
	}
}

func main() {
	var a checkfields = checkfields{
		m: 5,
		n: "hello world",
		l: 4.3,
	}
	check(a.m)
	check(a.n)
	check(a.l)
}
