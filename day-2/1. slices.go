// Go Slices

// Slices are similar to arrays, but are more powerful and flexible.

// Like arrays, slices are also used to store multiple values of the same type in a single variable.

// However, unlike arrays, the length of a slice can grow and shrink as you see fit.

// In Go, there are several ways to create a slice:

//     Using the []datatype{values} format
//     Create a slice from an array
//     Using the make() function

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	array := [...]int{4, 5, 6, 7}
	fmt.Printf("Type of elemenent is %T\n", slice)
	fmt.Printf("Type of elemenent is %T\n", array)

	// Create slice from array
	slice_from_arr := array[0:4]
	fmt.Println(slice_from_arr)

	// Create a Slice With The make() Function
	slice_with_make := make([]int, 5, 10)
	fmt.Printf("slice_with_make = %v\n", slice_with_make)
	fmt.Printf("len = %d\n", len(slice_with_make))
	fmt.Printf("cap = %d\n", cap(slice_with_make))

	// Printing by index
	fmt.Printf("first element of slice is %d\n", slice_with_make[0])
	new_val := 15
	fmt.Printf("I change %d to %d\n", slice_with_make[0], new_val)
	slice_with_make[0] = new_val
	fmt.Printf("Result %v\n", slice_with_make)

	// Adding elements
	value_to_add := 40
	fmt.Printf("I want to add %d to slice\n", value_to_add)
	fmt.Printf("Before %d\n", slice_with_make)
	slice_with_make = append(slice_with_make, value_to_add)
	fmt.Printf("After %d\n", slice_with_make)

	// Adding 2 slices
	fmt.Println("I want to add 2 arrays")
	merged_slice := append(slice_from_arr, slice_with_make...)
	fmt.Printf("Result %d\n", merged_slice)

	// Copying elements
	elemenent_to_copy := merged_slice[1:6]
	variable_to_copy := make([]int, len(elemenent_to_copy))
	copy(variable_to_copy, elemenent_to_copy)
	fmt.Printf("Elements are copied %d\n", variable_to_copy)
}
