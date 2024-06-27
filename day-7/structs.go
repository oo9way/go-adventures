package main

import "fmt"

type Address struct {
	country string
	city    string
}

type Person struct {
	firstName  string
	lastName   string
	middleName string
	age        int
	balance    int
}

func (p Person) getBornYear() int {
	return 2024 - p.age
}

func (p Person) getBalanceInUSD() int {
	return p.balance * 12600
}

func (p Person) getFullName() string {
	return fmt.Sprintf("%s %s %s", p.firstName, p.lastName, p.middleName)
}

func (p Person) getFormattedData() string {
	return fmt.Sprintf(
		"Person: %s\nAge: %d, was born in %d\nBalance in UZS: %d\nBalance in USD: %d",
		p.getFullName(), p.age, p.getBornYear(), p.balance, p.getBalanceInUSD())
}

func (a Address) string() string {
	return fmt.Sprintf("%s, %s", a.country, a.city)
}

func main() {
	address := Address{"Uzbekistan", "Tashkent"}
	fmt.Println(address.string())
	var address2 Address
	address2.country = "Palestine"
	address2.city = "Gaza"
	fmt.Println(address2.string())
	person := Person{firstName: "Joseph", lastName: "Rahimov", middleName: "Yusuf", age: 20, balance: 15000}
	fmt.Println(person.getFormattedData())
}
