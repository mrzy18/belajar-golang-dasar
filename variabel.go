package main

import "fmt"

func main() {
	// Deklarasi variabel manifest typing
	var firstName string = "John"
	// Deklarasi variabel type inference
	var lastName = "Doe"
	// Deklarasi variabel dengan keyword new
	var age *int = new(int)
	// Deklarasi variabel _
	var _ = "variabel ini tidak akan digunakan"

	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
	fmt.Println(age) // Menampilkan alamat memori dari variabel age
}
