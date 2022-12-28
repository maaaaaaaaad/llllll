package lib

import "fmt"

func init() {
	fmt.Println("Start Change number from pointer")
}

func ChangeNumber(number *int, upgrade int) {
	*number += upgrade
}
