package main

import (
	"fmt"
	"math"
)

func main() {
	printMessage("Hello world")
	var result int = sumNumber(5, 5)
	fmt.Println(result)
	divideNumber(10, 5)
	var area, circumference = calculate(10)
	fmt.Printf("Area: %.2f, Circumference: %.2f \n", area, circumference)
	var avg = average(2, 3, 5, 5, 8, 9, 7)
	var message = fmt.Sprintf("Average : %.2f", avg)
	fmt.Println(message)
}

func printMessage(message string) {
	fmt.Println(message)
}

// function with return
func sumNumber(a int, b int) int {
	return a + b
}

func divideNumber(a, b int) {
	if b == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %d\n", a, b)
		return
	}
	result := a / b
	fmt.Printf("%d : %d = %d\n", a, b, result)
}

// function multiple return
func calculate(d float64) (float64, float64) {
	area := math.Pi * math.Pow(d/2, 2)
	circumference := math.Pi * d

	return area, circumference
}

// function variadic
func average(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}
