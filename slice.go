package main

import "fmt"

func main() {
	var colors = []string{"red", "blue", "green", "black", "yellow"}
	fmt.Println(colors[0])
	fmt.Println(colors[0:2])             // [red, blue]
	fmt.Println(colors[0:4])             // [red, blue, green, black]
	fmt.Println(colors[:])               // [red, blue, green, black, yellow]
	fmt.Println(colors[3:])              // [black, yellow]
	fmt.Println(len(colors))             // 5
	fmt.Println(cap(colors))             // 5
	fmt.Println(append(colors, "brown")) // [red, blue, green, black, yellow, brown]
}
