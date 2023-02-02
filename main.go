package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("name: %s, age: %d", s.Name, s.Age)
}

func main() {
	student := Student{"mad", 22}
	var stringer Stringer
	stringer = student
	result := stringer.String()
	fmt.Println(result)
}
