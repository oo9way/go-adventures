package main

import "fmt"

func main() {
    // TODO
    printHelloWorld()
    printPersonalInfo("Oogway", "Master", 149, "+998900000001")
}

func printHelloWorld() {
    fmt.Println("Hello world with function")
}

func printPersonalInfo(firstName string, lastName string, age int, phone string){
    fmt.Printf("First name:\t%v\n", firstName)
    fmt.Printf("Last name:\t%v\n", lastName)
    fmt.Printf("Born:\t%v\n", calculateBirthYear(age))
    fmt.Printf("Age:\t%v\n", age)
    fmt.Printf("Phone:\t%v\n", phone)
}

func calculateBirthYear(age int) int {
    currentYear := 2024
    return currentYear - age
}
