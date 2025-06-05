package main

import (
	"fmt"
	"os"
)

func logging() {
	fmt.Println("selesai menjalankan app")
}

func runApp() {
	defer logging() // akan selalu dieksekusi terakhir walaupun terjadi error
	fmt.Println("Menjalankan app")
}

func main() {
	runApp()
	os.Exit(1)
	fmt.Println("blok kode lain") //tidak akan dieksekusi
}
