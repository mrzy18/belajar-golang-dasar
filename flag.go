package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int64("age", 18, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
}
