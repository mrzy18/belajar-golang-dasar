package main

import "belajar-golang-dasar/library"

func main() {
	library.SayHello()            // tidak error karena exported
	library.introduce("John doe") // error karena unexported
}
