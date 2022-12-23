package main

import (
	"fmt"
)

func main() {
	arr, i := []int{}, 0+1
	for i <= 100 {
		arr = append(arr, i)
		i++
	}
  result := 0
  for index, _ := range arr {
    result += arr[index]
  }
  fmt.Println(result)
}
