package main

import (
	"fmt"
	"os"
)

func main() {
	var argsRaw = os.Args
	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)
}
