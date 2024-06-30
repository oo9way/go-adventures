package main

import "fmt"

func errorHandler() {
	if r := recover(); r != nil {
		fmt.Println("Recovered ", r)
	}
}

func Divider(nominator int, divider int) float32 {
	defer errorHandler()
	if divider == 0 {
		panic("can't divide by 0")
	}
	return float32(nominator) / float32(divider)
}

func main() {
	no := Divider(10, 0)
	fmt.Println(no)

	no = Divider(10, 1)
	fmt.Println(no)
}
