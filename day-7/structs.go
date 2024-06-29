package main

import "fmt"

type Address struct {
	country string
	city    string
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
}
