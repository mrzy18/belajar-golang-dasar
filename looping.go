package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("looping:", i)
	}

	// for - range
	var students = [3]string{"john", "jane", "budi"}
	for _, student := range students {
		fmt.Println("student's name", student)
	}

	// for break and continue
	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("number: ", i)
	}

	// nested looping
	for i := 0; i <= 5; i++ {
		for j := i; j <= 5; j++ {
			fmt.Print(j, "")
		}
		fmt.Println()
	}
}
