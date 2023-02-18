package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()
	el1 := v.PushBack(1)
	el2 := v.PushFront(2)
	v.InsertAfter(3, el1)
	v.InsertBefore(5, el2)
	fmt.Println(v)

  for e := v.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value, " ")
  }

  for e := v.Back(); e != nil; e = e.Next() {
    fmt.Println(e.Value, " ")
  }
}
