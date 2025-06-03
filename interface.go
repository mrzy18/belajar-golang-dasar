package main

import (
	"fmt"
	"math"
)

type calc interface {
	area() float64
	circumference() float64
}
type circle struct {
	diameter float64
}

func (c circle) radius() float64 {
	return c.diameter / 2
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius(), 2)
}
func (c circle) circumference() float64 {
	return math.Pi * c.diameter
}

func main() {
	var shape calc

	shape = circle{10.0}
	fmt.Println("Area :", shape.area())
	fmt.Println("Circumference :", shape.circumference())
	fmt.Println("Radius :", shape.(circle).radius())
}
