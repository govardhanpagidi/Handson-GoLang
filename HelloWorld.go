package main

import (
	"fmt"
)

//*****How to run
// Open the terminal or cmd prompt, CD to helloworld.go dir, Run the below command
// go run HelloWorld.go

//alias print method
var print = fmt.Println

func main() {
	display("hello world!")
}

func display(text string){
	print(text)
}