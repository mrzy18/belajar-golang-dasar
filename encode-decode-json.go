package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// Decode
	jsonstring := `{"Name": "John Doe", "Age": 24}`
	jsonData := []byte(jsonstring)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Decode")
	fmt.Println("user :", data.FullName)
	fmt.Println("Age :", data.Age)

	// Encode
	var object = []User{{"john doe", 27}, {"Richard Dennis", 32}}
	var jsonObject, e = json.Marshal(object)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	var result = string(jsonObject)
	fmt.Println("Encode")
	fmt.Println(result)
}
