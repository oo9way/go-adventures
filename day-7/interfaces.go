package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rectangle struct {
	x int
	y int
}

func (r Rectangle) area() int {
	return r.x * r.y
}

func (r Rectangle) location() Point {
	return Point{x: r.x, y: r.y}
}

func main() {
	var rectangle Rectangle = Rectangle{x: 10, y: 15}
	fmt.Println(rectangle.area())
}
