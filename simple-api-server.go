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

var data = []Product{
	Product{"P001", "handphone", "$200"},
	Product{"P002", "laptop", "$500"},
	Product{"P003", "keyboard", "$120"},
	Product{"P004", "camera", "$450"},
}

func products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "GET" {
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func product(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id := r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.Id == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/products", products)
	http.HandleFunc("/product", product)

	fmt.Println("server is running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
