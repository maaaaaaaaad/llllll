package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func main() {
	var a *account = &account{100, "mad", "integral"}
	fmt.Println("original a:", a)
	a.withdrawPointer(20)
	fmt.Println("new a:", a)
}
