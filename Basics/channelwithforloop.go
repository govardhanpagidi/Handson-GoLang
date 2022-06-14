// You can edit this code!
// Click here and start typing.

//This is helpful when you want to process bigger array of elements without memory issues
//Uses channel to process one by one rather than holding the entire array in memory
package main

import (
	"fmt"
	"time"
)

func main() {

	primChan := make(chan int,16)
	childChan := make(chan int,16)

	i := 0
	go func() {
		for i = 0; i < 10000; i++ {
			primChan <- i
			fmt.Println("pushed  to parent channel", i)
		}
		close(primChan)

	}()

	go func(){
		for num := range primChan {
		time.Sleep(1 * time.Second)
		childChan <- num
			fmt.Println("pushed  to child channel", num)
		}
		close(childChan)
	}()

	for num := range childChan{
		go func(){
		time.Sleep(2 * time.Second)
		fmt.Println("num received ", num)
		}()
	}
}
