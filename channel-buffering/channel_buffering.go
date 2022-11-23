package main

import "fmt"

func main() {
	messages := make(chan string, 4)

	messages <- "Buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
