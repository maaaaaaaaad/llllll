package main

import (
	"container/list"
	"fmt"
)


func main() {
	v := list.New()
  e1 := v.PushFront(1)
  e2 := v.PushBack(2)
  v.InsertAfter(3, e1)
  v.InsertAfter(4, e2)

  for e := v.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value, "")
  }
}
