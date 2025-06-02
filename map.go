package main

import (
	"fmt"
)

func main() {
	var person map[string]string
	person = map[string]string{}
	person["name"] = "john doe"
	person["age"] = "24"
	person["date of birth"] = "18/07/2000"

	var product = map[string]string{"product": "phone", "price": "$400", "quantity": "10"}
	fmt.Println(product)
	for key, val := range product {
		fmt.Println(key, "  \t:", val)
	}

	var chikens = []map[string]string{
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken green", "gender": "female"},
		map[string]string{"name": "chicken blue", "gender": "female"},
	}
	for _, chiken := range chikens {
		fmt.Println(chiken["name"], chiken["gender"])
	}
}
