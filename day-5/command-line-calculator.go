package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, err := strconv.Atoi(os.Args[1])
	b, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println(err)
		fmt.Println("Only integer digits")
	} else {
		fmt.Println(a + b)
	}
}
