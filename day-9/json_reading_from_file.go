package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Products struct {
	Items []Product `json:"products"`
}

func main() {
	file, err := ioutil.ReadFile("products.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	data := Products{}

	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for i := 0; i < len(data.Items); i++ {
		fmt.Println("Product ID:", data.Items[i].Id)
		fmt.Println("Name:", data.Items[i].Name)
	}
}
