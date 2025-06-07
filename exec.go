package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var output1, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf(" -> git config user.name\n%s\n", string(output1))

	var output2, _ = exec.Command("go", "version").Output()
	fmt.Printf(" -> go version\n%s\n", string(output2))
}
