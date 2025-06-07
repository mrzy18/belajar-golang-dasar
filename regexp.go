package main

import (
	"fmt"
	"regexp"
)

func main() {
	var str = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Method ini mengembalikan nilai boolean, jika string sesuai dengan regexp akan bernilai true
	fmt.Println(regex.MatchString(str))
	// Method ini akan memisahkan substring yang memenuhi kriteria regexp
	fmt.Println(regex.Split(str, -1))
	// Method untuk mencari string yang sesuai kriteria regexp
	fmt.Println(regex.FindString(str))
	// Method untuk mencari string yang sesuai kriteria regexp tapi yang dikembalikan index-nya
	fmt.Println(regex.FindStringIndex(str))
	// Method untuk mencari banyak setring yang sesuai kriteria regexp
	fmt.Println(regex.FindAllString(str, -1))
	// Method untuk me-replace semua string yang sesuai kriteria regexp
	fmt.Println(regex.ReplaceAllString(str, "potato"))
	// Sama seperti method sebelumnya tapi menggunakan function
	fmt.Println(regex.ReplaceAllStringFunc(str, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	}))
}
