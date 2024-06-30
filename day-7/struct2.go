package main

import "fmt"

type Product struct {
	title        string
	description  string
	quantity     int
	pricePerUnit string
	total        string
}

func (p Product) string() string {
	return fmt.Sprintf("Title: %s\nDescription: %s\nQuantity: %v\nPrice per unit: %s\nTotal: %s\n", p.title, p.description, p.quantity, p.pricePerUnit, p.total)
}

func main() {
	var product Product
	product.title = "Go Programming Language"
	product.description = "The Go Programming Language"
	product.quantity = 12
	product.pricePerUnit = "12$"
	product.total = "144$"
	fmt.Println(product.string())
}
