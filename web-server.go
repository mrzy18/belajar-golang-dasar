package main

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	http.HandleFunc("/ping", pong)

	fmt.Println("server is running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
