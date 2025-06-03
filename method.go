package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	age  uint8
}

func (P Person) sayHello() {
	fmt.Println("halo", P.name)
}

func (P Person) getNameAt(i int) string {
	return strings.Split(P.name, " ")[i-1]
}

// Method Pointer
func (P *Person) changeName(newName string) {
	fmt.Println(P.name, "changed to", newName)
	P.name = newName
}

func main() {
	var johnDoe = Person{"John Doe", 25}
	johnDoe.sayHello()
	var lastName = johnDoe.getNameAt(2)
	fmt.Println("last name:", lastName)
	johnDoe.changeName("John Wick")
	fmt.Println(johnDoe.name)
}
