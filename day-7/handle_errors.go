package main

import (
	"errors"
	"fmt"
)

var NoTooSmall = errors.New("Too small")

func checkNumber(n int) (int, error) {
	if n > 0 {
		return n, nil
	} else {
		return 0, NoTooSmall
	}
}

func main() {
	n, err := checkNumber(-10)
	fmt.Println(n, err)
}
