package main

import (
	"fmt"
	"runtime"
)

func printString(looping int, message string) {
	for i := 0; i < looping; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go printString(5, "goroutine")
	printString(5, "not goroutine")

	var input string
	fmt.Scanln(&input)
}
