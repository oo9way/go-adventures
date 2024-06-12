package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Default value ", i)

	zeroval(i)
	fmt.Println("After zeroval", i)

	zeroptr(&i)
	fmt.Println("After zeroptr", i)

	fmt.Println("Pointer ", &i)
}
