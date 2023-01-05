package main

import (
	"fmt"
)

type person struct {
  name string
  age int
}

func (receiver *person) changer(name string , age int) {
  receiver.name =  name
  receiver.age = age
}

func main() {
 var human *person = new(person)
  human.name = "mad"
  human.age = 22

  fmt.Println(human)

  human.changer("in", 20)

  fmt.Println(human)
}
