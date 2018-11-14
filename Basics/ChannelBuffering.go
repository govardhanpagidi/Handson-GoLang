package main

import "fmt"

func main(){

	messages := make(chan string, 2)

	messages <- "Pagidi"
	messages <- "Govardhan"

	fmt.Println( <- messages)
	fmt.Println( <- messages)
}