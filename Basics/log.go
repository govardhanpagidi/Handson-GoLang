package main

import "fmt"
import "os"
import "log"

func init(){
	nf,err := os.Create("log.txt")
	if err != nil{
		fmt.Print(err)
	}
	log.SetOutput(nf)
}


func main(){

	_,err := os.Open("abs.txt")

	if err != nil{
		log.Println("err happened",err)
		log.Fatal(err)
		panic(err)
	}
}