package main

import (
	"fmt"
)

func Divide(nominator int, divider int) float32 {
	if divider == 0 {
		panic("can't divide by 0")
	}
	return float32(nominator) / float32(divider)
}

func main() {
	no := Divide(10, 0)
	fmt.Println(no)
	no = Divide(10, 1)
	fmt.Println(no)
}
