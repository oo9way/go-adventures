package main

import "fmt"

func main() {
	var command string
	contacts := make(map[string]string)
	fmt.Println("Welcome to phone book")

	for {
		fmt.Println("Commands\n  1.add - to add a phone\n  2.list - list of phones\n  3.lookup - find with name\n  4.exit - Exit from phonebook\n")
		fmt.Print("Enter a command: ")
		fmt.Scan(&command)
		if command == "add" {
			fmt.Print("Enter a name and a phone with space between: ")
			var phone string
			var name string
			fmt.Scan(&name, &phone)
			contacts[name] = phone
			fmt.Println("Contact is saved")
		} else if command == "list" {
			fmt.Println("List of phones in Phone Book")
			for key, value := range contacts {
				fmt.Println(key, value)
			}
			fmt.Println("===============")
		} else if command == "lookup" {
			fmt.Print("Enter a name of contact: ")
			var name string
			fmt.Scan(&name)
			fmt.Println(contacts[name])
		} else if command == "exit" {
			break
		} else {
			fmt.Println("Invalid command ", command)
		}
	}
	fmt.Println("Exiting ...")
}
