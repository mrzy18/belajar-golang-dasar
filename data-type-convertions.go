package main

import (
	"fmt"
	"strconv"
)

func main() {
	strToInt, _ := strconv.Atoi("123")                  // konversi string ke int
	strParseInt, _ := strconv.ParseInt("123", 10, 64)   // konversi string numerik ke numerik non-desimal
	strParseFloat, _ := strconv.ParseFloat("24.12", 32) // konversi string ke numerik desimal
	strParseBool, _ := strconv.ParseBool("true")        // konversi string ke boolean
	fmt.Println(strToInt)
	fmt.Println(strParseInt)
	fmt.Println(strParseFloat)
	fmt.Println(strParseBool)
	fmt.Println(strconv.Itoa(123))                               // konversi int ke string
	fmt.Println(strconv.FormatInt(int64(24), 8))                 // konversi numerik int64 ke string
	fmt.Println(strconv.FormatFloat(float64(24.12), 'f', 6, 64)) // konversi float64 ke string
	fmt.Println(strconv.FormatBool(true))                        // konversi boolean ke string

	// konversi juga bisa menggunakan teknik casting
	fmt.Println(float64(24))
	fmt.Println(int32(24.00))

}
