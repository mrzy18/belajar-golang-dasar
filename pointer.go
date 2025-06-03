package main

import "fmt"

func main() {
	// secara default go akan menduplikat isi dari variabel
	var numberA int = 5
	var numberB int = numberA
	numberB = 10

	// valuenya berbeda karena merujuk pada alamat memori yang berbeda
	fmt.Println(numberA) // 5
	fmt.Println(numberB) // 10

	var name string = "John Doe"
	var newName *string = &name
	*newName = "Jeremy Burns"

	// valuenya sama karena merujuk pada alamat memori yang sama
	fmt.Println(name)     // Jeremy Burns
	fmt.Println(*newName) // Jeremy Burns

	var value int = 3
	fmt.Println("value before change: ", value)
	changeValue(&value, 100)
	fmt.Println("value after change: ", value)
}

func changeValue(originalValue *int, newValue int) {
	*originalValue = newValue
}
