// Declare an Array

// In Go, there are two ways to declare an array:

// 1. With the var keyword:
// Syntax
// 		var array_name = [length]datatype{values} // here length is defined
// 		var array_name = [...]datatype{values} // here length is inferred

// 2. With the := sign:
// Syntax
// 		array_name := [length]datatype{values} // here length is defined
// 		array_name := [...]datatype{values} // here length is inferred

package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	var arr3 = [...]int{10, 11, 12}
	arr4 := [...]int{13, 14, 15, 16, 17, 18}

	var cars = [...]string{"BMW", "Volvo", "GM"}
	var students = [...]string{"Shifu", "Panda", "Tiger", "Monkey", "Snake"}

	null_arr := [5]int{}
	var partial_arr = [5]int{1, 2, 3}
	full_arr := [5]int{1, 2, 3, 4, 5}

	var specific_index = [5]int{2: 10, 4: 20}

	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(arr3, len(arr3))
	fmt.Println(arr4, len(arr4))

	fmt.Println(cars)

	cars[2] = "Chevrolet"

	fmt.Printf("%v rides %v\n", students[0], cars[0])
	fmt.Printf("%v rides %v\n", students[1], cars[1])
	fmt.Printf("%v rides %v\n", students[2], cars[2])

	fmt.Printf("%v null\n", null_arr)
	fmt.Printf("%v partial\n", partial_arr)
	fmt.Printf("%v full\n", full_arr)

	fmt.Printf("%v with specific index", specific_index)

}
