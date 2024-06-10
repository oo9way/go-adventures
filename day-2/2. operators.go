package main

import "fmt"

func main() {
	var a = 15 + 25
	var sum1 = 100 + 150
	var sum2 = 100 + sum1
	var sum3 = sum1 + sum2

	// Arithmethics + - * / % ++ --
	// Assignment = += -= *= /= %= &= |= ^= >>= <<=
	// Comparison == != < > >= <=
	// Logical && || !
	// Bitwise (only binary) & | ^ << >>

	fmt.Println(a, sum1, sum2, sum3)
}
