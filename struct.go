package main

import "fmt"

func main() {
	type Person struct {
		name       string
		age        uint8
		gender     string
		isMarriage bool
	}

	var johnDoe = Person{"John Doe", 25, "male", false}
	fmt.Println("person: ", johnDoe)

	// Embedded Struct
	type Student struct {
		grade uint8
		Person
	}
	var student = Student{}
	student.Person.name = "Philip Patton"
	student.Person.age = 22
	student.grade = 2
	student.Person.gender = "male"
	student.Person.isMarriage = false
	fmt.Println("student: ", student)

	// kombinasi struct dan slice
	var people = []Person{
		{"John Doe", 27, "male", true},
		{"Brandon Malone", 21, "male", false},
		{"Jennie Poole", 23, "female", false},
	}

	for _, person := range people {
		fmt.Println(person.name, "is", person.age, "years old")
	}
}
