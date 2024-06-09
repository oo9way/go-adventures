// The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
// const CONSTNAME type = value
// Multiple constants can be grouped together into a block for readability:

package main

import "fmt"

const PI = 3.14
const A int = 1
const B = "hello world"

const (
	X int     = 1
	Y float32 = 4.5
	Z string  = "Hello world"
)

func main() {
	fmt.Println(PI, A)
	fmt.Printf("X - %g\nY - %g\nZ - %g", X, Y, Z)
}
