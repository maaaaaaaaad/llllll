package main

import (
	"fmt"
)

func main() {
	msg := "Hello"
	var saying = say(msg)
	fmt.Println(saying)
}

func say(msg string) string {
	return msg
}
