package main

import "fmt"

func Sum[V int | float64](numbers []V) V {
	var total V
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total: ", result)
}
