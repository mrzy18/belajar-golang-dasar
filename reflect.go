package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("Variable type :", reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Variable value :", reflectValue.Int())
	}
}
