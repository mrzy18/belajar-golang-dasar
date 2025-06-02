package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "banana", "grape", "melon"}
	fmt.Println("array length \t", len(fruits))
	fmt.Println("array values \t", fruits)
	// inisialisasi array gaya vertikal
	var students = [5]string{
		"john",
		"jane",
		"billy",
		"budi",
		"evan",
	}
	fmt.Println("students:", students)
	//inisialisasi array tanpa jumlah element
	var numbers = [...]uint8{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	//array multidimenssi
	var numbers1 = [2][3]uint8{[3]uint8{1, 2, 3}, [3]uint8{1, 2, 3}}
	var numbers2 = [2][3]uint8{{123}, {123}}
	fmt.Println(numbers1)
	fmt.Println(numbers2)
	//membuat array dengan keyword make
	var colors = make([]string, 2)
	colors[0] = "red"
	colors[1] = "black"
	fmt.Println(colors)
}
