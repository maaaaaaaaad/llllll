package main

import (
	"fmt"
	"main/lib"
)

var num int = 1000

func init() {
	fmt.Println("Start inspect your number:", num)
}

func main() {
	result := lib.InspectNumberThan100(num)
	fmt.Println(result)
}
