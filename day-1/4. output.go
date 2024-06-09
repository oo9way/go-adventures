// Go has three functions to output text:

// Print()
// Println()
// Printf()

// The Print() function prints its arguments with their default format.
// The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end

// The Printf() function first formats its argument based on the given formatting verb and then prints them.
// Here we will use two formatting verbs:

//     %v is used to print the value of the arguments
//     %T is used to print the type of the arguments

// \n creates a new line

// ========================================================== //
// General Formatting Verbs

// The following verbs can be used with all data types:
// 	Verb 	Description
// 	%v 		Prints the value in the default format
// 	%#v 	Prints the value in Go-syntax format
// 	%T 		Prints the type of the value
// 	%% 		Prints the % sign

// ========================================================== //

// Integer Formatting Verbs

// The following verbs can be used with the integer data type:
// 	Verb 	Description
// 	%b 		Base 2
// 	%d 		Base 10
// 	%+d 	Base 10 and always show sign
// 	%o 		Base 8
// 	%O 		Base 8, with leading 0o
// 	%x 		Base 16, lowercase
// 	%X 		Base 16, uppercase
// 	%#x 	Base 16, with leading 0x
// 	%4d 	Pad with spaces (width 4, right justified)
// 	%-4d 	Pad with spaces (width 4, left justified)
// 	%04d 	Pad with zeroes (width 4

// ========================================================== //

// String Formatting Verbs

// The following verbs can be used with the string data type

// 	Verb 	Description
// 	%s 		Prints the value as plain string
// 	%q 		Prints the value as a double-quoted string
// 	%8s 	Prints the value as plain string (width 8, right justified)
// 	%-8s 	Prints the value as plain string (width 8, left justified)
// 	%x 		Prints the value as hex dump of byte values
// 	% x 	Prints the value as hex dump with spaces

// ========================================================== //

// Boolean Formatting Verbs

// The following verb can be used with the boolean data type:
// 	Verb 	Description
// 	%t 		Value of the boolean operator in true or false format (same as using %v)

// ========================================================== //

// Float Formatting Verbs

// The following verbs can be used with the float data type:
// 	Verb 	Description
// 	%e 		Scientific notation with 'e' as exponent
// 	%f 		Decimal point, no exponent
// 	%.2f 	Default width, precision 2
// 	%6.2f 	Width 6, precision 2
// 	%g 		Exponent as needed, only necessary digits

package main

import (
	"fmt"
)

func main() {
	var username, password string = "Oogway", "*****"
	var firstName, lastName string = "Master", "Oogway"
	var age, birthDate int = 149, 1875

	fmt.Print("Username: ", username, "\n")
	fmt.Print("Password: ", password, "\n")
	fmt.Println("=================")
	fmt.Println("First name", firstName) // We do not need to define \n to write from new line
	fmt.Println("Last name", lastName)
	fmt.Printf("was born in %v, %v years old \n", birthDate, age)
	fmt.Printf("variable age is %T, birthDate is %T", age, birthDate)
}
