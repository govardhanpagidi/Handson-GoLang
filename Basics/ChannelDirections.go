package main

import(
	"fmt"
)

func main(){

	pings := make(chan string, 1)
	pongs := make(chan string,1)

	ping(pings,"Hi")
	pong(pings,pongs)

	fmt.Println(<-pongs)
}

func ping(pings chan<- string,msg string){
	fmt.Println("acknowledging.. ",msg)
	pings <- msg
}

func pong(pings <-chan string,pongs chan<- string){
	fmt.Println("Received ping")
	msg := <- pings
	pongs <- msg
}

