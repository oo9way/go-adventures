package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var person1 Person
	var person2 Person

	// Oogway
	person1.name = "Oogway"
	person1.age = 149
	person1.job = "Master"
	person1.salary = 500

	// Shifu
	person2.name = "Shifu"
	person2.age = 100
	person2.job = "Master"
	person2.salary = 100

	prettyPrintPerson(person1)
	prettyPrintPerson(person2)

}

func prettyPrintPerson(person Person) {
	fmt.Printf("Name:\t%v\n", person.name)
	fmt.Printf("Age:\t%v\n", person.age)
	fmt.Printf("Job:\t%v\n", person.job)
	fmt.Printf("Salary:\t%v\n", person.salary)

	fmt.Println("=============")
}
