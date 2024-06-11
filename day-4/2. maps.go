package main

import "fmt"

func main() {
	salaries := map[string]int{"Oogway": 500, "Shifu": 100, "Panda": 50, "Tiger": 20, "Monkey": 0}
	salaries["Snake"] = 0
	salaries["Bear"] = 200

	for salary := range salaries {
		delete(salaries, "Bear")
		fmt.Printf("%v\t%v\n", salary, salaries[salary])
	}
	salary, inList := salaries["Bear"]
	fmt.Printf("Is a Bear in the list? %v %v\n", inList, salary)
}
