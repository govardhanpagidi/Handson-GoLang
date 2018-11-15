package main

import (
	"fmt"
	// "time"
	"sync"
)

var wg sync.WaitGroup 

func main(){

	for i:= 0; i< 5; i++{
		wg.Add(1)
		go routine(i)
	}
	wg.Wait()
	fmt.Print("Main thread is closed.")
}

func routine(i int){

	fmt.Println("Routine finished : " ,i)
	defer wg.Done()
}
