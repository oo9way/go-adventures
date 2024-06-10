package main

import "fmt"

func main() {
	day := 10

	day_number := 5

	switch day {
	case 1:
		fmt.Println("It's a Monday")
	case 2:
		fmt.Println("It's a Tuesday")
	case 3:
		fmt.Println("It's a Wendsday")
	case 4:
		fmt.Println("It's a Thursday")
	case 5:
		fmt.Println("It's a Friday")
	case 6:
		fmt.Println("It's a Sunday")
	case 7:
		fmt.Println("It's a Saturday")
	default:
		fmt.Println("No weekdays")
	}

	switch day_number {
	case 1, 3, 5, 7:
		fmt.Println("Odd day")
	case 2, 4, 6:
		fmt.Println("Even day")
	}
}
