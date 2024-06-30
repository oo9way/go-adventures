package main

import "fmt"

func main() {
	arr := make([]string, 0)
	var response string

	for {
		fmt.Print("Enter a command > ")
		fmt.Scan(&response)

		if response == "new" {
			fmt.Println("Enter a new entry")
			fmt.Scan(&response)
			arr = append(arr, response)
			fmt.Println("New entry is saved")
		} else if response == "list" {
			fmt.Println("List of entries")
			if len(arr) == 0 {
				fmt.Println("There are no entries")
			} else {
				for i := 0; i < len(arr); i++ {
					fmt.Printf("%v. %v\n", i+1, arr[i])
				}
			}
		} else if response == "exit" {
			fmt.Println("Exiting ...")
			break
		} else {
			fmt.Println("Invalid command")
		}
	}

}
