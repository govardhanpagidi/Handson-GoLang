package main

import "fmt"
import "os"
import "log"

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Print(err)
	}
	log.SetOutput(nf)
}

func main() {

	// Open a file
	log.Println("Trying to open a file.")

	defer fileOp()

	log.Println("After opening a file")
}

func fileOp() {
	_, err := os.Open("abs.txt")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
