package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

func main() {
	var data = student{
		name:        "John doe",
		height:      182.5,
		age:         26,
		isGraduated: false,
		hobbies:     []string{"eating", "sleeping", "coding"},
	}
	fmt.Printf("%b\n", data.age)           // data numerik -> string numerik biner 26 -> 11010
	fmt.Printf("%c\n", 97)                 // data numerik(unicode) -> string karakter unicode 97 -> a
	fmt.Printf("%d\n", data.age)           // data numerik -> string numerik desimal 26 -> "26"
	fmt.Printf("%e\n", data.height)        // data numerik desimal -> scientific notation 182.5 -> 1.825000e+02
	fmt.Printf("%f\n", data.height)        // 182.5 -> 182.500000
	fmt.Printf("%.2f\n", data.height)      // 182.5 -> 182.50
	fmt.Printf("%o\n", data.age)           // data numerik -> string numerik oktal 26 -> "32"
	fmt.Printf("%p\n", data.name)          // pointer 0x2081be0c0
	fmt.Printf("%q\n", `" name \ height"`) // escape string
	fmt.Printf("%s\n", data.name)          // untuk format string (John doe)
	fmt.Printf("%t\n", data.isGraduated)   // untuk format boolean (false)
	fmt.Printf("%T\n", data.name)          // untuk format type data (string)
	fmt.Printf("%v\n", data)               // untuk format data apa saja
	fmt.Printf("%+v\n", data)              // untuk format struct
	fmt.Printf("%#v\n", data)              // mengembalikan nama dan nilai tiap property sesuai dengan struktur struct
	fmt.Printf("%x\n", data.age)           // data numerik -> string numerik hexadesimal
	fmt.Printf("%%\n")                     // menulis karakter % pada format string
}
