package main

import (
	"errors"
	"fmt"
	"strings"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error Occured")
	} else {
		fmt.Println("Application running perfectly")
	}
}

func isPalindrome(word string) (bool, error) {
	if strings.TrimSpace(word) == "" {
		return false, errors.New("input cannot be empty")
	}
	word = strings.ToLower(word)
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false, errors.New("The word is not palindrome")
		}
	}
	return true, nil
}

func main() {
	defer catch()
	var word string
	fmt.Print("Type a word: ")
	fmt.Scanln(&word)

	if valid, err := isPalindrome(word); valid {
		fmt.Printf("The word %s is palindrome \n", word)
	} else {
		panic(err.Error())
	}
}
