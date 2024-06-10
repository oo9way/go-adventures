package main

import "fmt"

func main() {
	panda := 15
	shifu := 40
	oogway := 150

	if shifu < panda && oogway < panda {
		fmt.Println("Master is Panda")
	} else if panda < shifu && oogway < shifu {
		fmt.Println("Master is Shifu")
	} else {
		fmt.Println("Master is Oogway")
	}
}
