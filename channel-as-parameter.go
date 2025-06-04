package main

import "fmt"

// Channel as Parameter
func cetakPesan(pesan chan string) {
	fmt.Println(<-pesan)
}

func main() {
	var pesan = make(chan string)
	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			pesan <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		cetakPesan(pesan)
	}
}
