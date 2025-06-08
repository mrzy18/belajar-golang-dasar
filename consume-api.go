package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id          string `json:"id"`
	Productname string `json:"productname"`
	Price       string `json:"price"`
}

var baseURL = "http://localhost:3000"

func fetchProducts() ([]Product, error) {
	var err error
	var client = &http.Client{}
	var data []Product

	request, err := http.NewRequest("GET", baseURL+"/products", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	var products, err = fetchProducts()
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	for _, product := range products {
		fmt.Printf("id: %s\nproductName: %s\nprice: %s\n", product.Id, product.Productname, product.Price)
	}
}
