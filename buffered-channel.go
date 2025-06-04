package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2)
	go func() {
		for {
			i := <-messages
			fmt.Println("Received data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Send data", i)
		messages <- i
	}

	time.Sleep(1 * time.Second)
}
