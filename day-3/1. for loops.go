package main

import "fmt"

func main () {
    phones := [...]string{"iPhone 5", "iPhone 6", "Samsung",  "iPhone 7", "Android", "Xiaomi",  "iPhone 8", "iPhone 9", "iPhone X"}
    versions := [...]string{"S", "XS", "pro", "pro max"}

    // Basic loop
    for i:=0; i < len(phones); i++ {
        for j:=0; j < len(versions); j++ {
            if (phones[i] == "Samsung" || phones[i] == "Xiaomi" || phones[i] == "Android") {
                continue
            }
            fmt.Printf("%v %v\n", phones[i], versions[j])
        }
    }
    laptops := [...]string{"Macbook", "HP", "Asus", "Acer"}

    // With range
    for id, laptop := range laptops {
        fmt.Printf("%v\t%v\n", id, laptop)
    }

    fruits := [...]string{"banana", "strawberry", "apple", "cherry"}
    for _, fruit := range fruits {
        fmt.Printf("%v\n", fruit)
    }
}