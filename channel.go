package main

import "fmt"

func main() {
	var messages = make(chan string)
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("Tony Flowers")
	go sayHelloTo("Jonathan Colon")
	go sayHelloTo("Isabella Morton")

	var message1 = <-messages
	fmt.Println(message1)
	var message2 = <-messages
	fmt.Println(message2)
	var message3 = <-messages
	fmt.Println(message3)
}
