package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("john doe", "doe"))            // true
	fmt.Println(strings.HasPrefix("john doe", "jo"))            // true
	fmt.Println(strings.HasSuffix("john doe", "oe"))            // true
	fmt.Println(strings.Count("john doe", "o"))                 // 2
	fmt.Println(strings.Index("john doe", "d"))                 // 5
	fmt.Println(strings.Replace("john doe", "john", "jane", 1)) // jane doe
	fmt.Println(strings.Repeat("HA", 3))                        // HAHAHA
	fmt.Println(strings.Split("john doe", " "))                 // ["john", "doe"]
	fmt.Println(strings.Join([]string{"john", "doe"}, " "))     // john doe
	fmt.Println(strings.ToLower("HALLO"))                       // hallo
	fmt.Println(strings.ToUpper("eat!"))                        // EAT!
}
