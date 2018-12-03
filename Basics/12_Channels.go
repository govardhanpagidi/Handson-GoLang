package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		fmt.Println("anonymous")

	}()

	msg := <-messages

	fmt.Println("main thread")
}
