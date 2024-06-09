// Go has three basic data types:

//     bool: represents a boolean value and is either true or false
//     Numeric: represents integer types, floating point values, and complex types
//     string: represents a string value

// Boolean Data Type
// A boolean data type is declared with the bool keyword and can only take the values true or false.
// The default value of a boolean data type is false

// Integer Data Types
// Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.
// The integer data type has two categories:
//     Signed integers - can store both positive and negative values
//     Unsigned integers - can only store non-negative values

// Signed Integers
// Signed integers, declared with one of the int keywords, can store both positive and negative values:

// Unsigned Integers
// Unsigned integers, declared with one of the uint keywords, can only store non-negative values:

// Go has five keywords/types of signed integers:
// Type 	Size 			Range
// int 		Depends on platform: 32 bits in 32 bit systems and 64 bit in 64 bit systems -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems
// int8 	8 bits/1 byte 	-128 to 127
// int16 	16 bits/2 byte 	-32768 to 32767
// int32 	32 bits/4 byte 	-2147483648 to 2147483647
// int64 	64 bits/8 byte 	-9223372036854775808 to 9223372036854775807

// Go has five keywords/types of unsigned integers:
// Type 	Size 			Range
// uint 	Depends on platform: 32 bits in 32 bit systems and 64 bit in 64 bit systems 	0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems
// uint8 	8 bits/1 byte 	0 to 255
// uint16 	16 bits/2 byte 	0 to 65535
// uint32 	32 bits/4 byte 	0 to 4294967295
// uint64 	64 bits/8 byte 	0 to 18446744073709551615

// Go Float Data Types
// The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.

// The float data type has two keywords:
// Type 	Size 		Range
// float32 	32 bits 	-3.4e+38 to 3.4e+38.
// float64 	64 bits 	-1.7e+308 to +1.7e+308.

package main

import "fmt"

func main() {
	var x int = 100
	var y int = -100

	var a uint = 100
	var b uint = 200

	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	fmt.Printf("Type: %T, value: %v\n", a, a)
	fmt.Printf("Type: %T, value: %v\n", b, b)
}
