package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type OrderItem struct {
	Id       int     `json:"id"`
	Quantity int     `json:"quantity"`
	Total    float32 `json:"total"`
}

type Order struct {
	Id    int         `json:"id"`
	Items []OrderItem `json:"items"`
}

type Response struct {
	Orders []Order `json:"orders"`
}

func main() {
	file, _ := ioutil.ReadFile("orders.json")
	data := Response{}

	_ = json.Unmarshal([]byte(file), &data)
	fmt.Println(data)
}
