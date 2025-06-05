package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(name string) (bool, error) {
	if strings.TrimSpace(name) == "" {
		// custom error
		return false, errors.New("name cannot be empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("hello, ", name)
	} else {
		fmt.Println(err.Error())
	}
}
