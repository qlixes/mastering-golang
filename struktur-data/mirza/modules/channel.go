package modules

import "fmt"

func ChannelFunction() {
	// deklarasi channel
	channel := make(chan string)

	// goroutine
	go func() {
		channel <- "Hello, World!"
	}()

	// print channel value
	fmt.Println(<-channel)
}
