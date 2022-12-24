package lib

import "fmt"

func init() {
	fmt.Println("Inspecting int type more than 100")
}

func InspectNumberThan100(arg int) bool {
	compare := 100
	return arg > compare
}
