package main

import "fmt"

func main() {
	var point uint8 = 6
	switch {
	case point >= 8:
		fmt.Println("Perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		fmt.Println("you need to learn more!")
	}
}
