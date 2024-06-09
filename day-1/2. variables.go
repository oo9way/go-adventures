// Declaration variables
// var name type = expression
// short variable declarations :=, a := 1.5
// can't use short variable declarations outside of function
// it is possible to declare multiple variables in the same line.
// If the type keyword is not specified, you can declare different types of variables in the same line
// Multiple variable declarations can also be grouped together into a block for greater readability:

// Go variable naming rules:

//     A variable name must start with a letter or an underscore character (_)
//     A variable name cannot start with a digit
//     A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
//     Variable names are case-sensitive (age, Age and AGE are three different variables)
//     There is no limit on the length of the variable name
//     A variable name cannot contain spaces
//     The variable name cannot be any Go keywords

package main

import "fmt"

// Global variable
const boilingF = 212.0

// r := "outside function"   error comes

func main() {
	// Local variables
	var f = boilingF
	var c = (f - 32) * 5 / 9
	var d string
	var i, j, k, l int = 1, 2, 3, 4
	e := 10
	x, y := 4, "HOO"
	var (
		name    string = "Oogway"
		surname string = "Master"
		age     int    = 149
	)

	d = "Changing type"

	fmt.Printf("Boiling point = %g*F or %g*C\n", f, c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(i, j, k, l)
	fmt.Println(x, y)
	fmt.Println(name, surname, age)
}
