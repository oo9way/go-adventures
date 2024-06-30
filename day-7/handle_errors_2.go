package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func errorHandler() {
	if r := recover(); r != nil {
		log.Println(r, string(debug.Stack()))
	}
}

func Divider(nominator int, divider int) float32 {
	defer errorHandler()
	if divider == 0 {
		panic("Can't divide by 0")
	}
	return float32(nominator) / float32(divider)
}

func main() {
	f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)
	log.Println("Starting program")
	no := Divider(10, 0)
	fmt.Println(no)

	no = Divider(10, 1)
	fmt.Println(no)
	f.Close()
}
