package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()
	f := v.PushFront(1)
	b := v.PushBack(2)
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	v.InsertAfter(10, f)
	v.InsertBefore(20, b)
	fmt.Println()
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 1,10,20,2
	}
}
