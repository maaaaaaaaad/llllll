package main

import (
	"fmt"
	"main/lib"
)

func main() {
	num, changeNum := 2, 10
	fmt.Println("init num:", num)
	lib.ChangeNumber(&num, changeNum)
	fmt.Println("change from pointer num:", num)
}
