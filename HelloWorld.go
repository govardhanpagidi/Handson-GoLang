// package main

// import (
// 	"fmt"
// )

// //*****How to run
// // Open the terminal or cmd prompt, CD to helloworld.go dir, Run the below command
// // go run HelloWorld.go

// //alias print method
// var print = fmt.Println

// func main() {
// 	display("hello world!")
// }

// func display(text string){
// 	print(text)
// }

package main

import (
	"fmt"
	"time"
)

func f1(c chan int, quit chan bool) {
	i := 0
	ticker := time.NewTicker(5000 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			i++
			fmt.Println("Generated", i)
			c <- i
		case <-quit:
			ticker.Stop()
			fmt.Println("Ticker stopped. Closing channel.")
			//quit <- true
			close(c)
			//close(quit)
			return
		}
	}
}

func main() {
	c := make(chan int)
	q := make(chan bool)

	go f1(c, q)

	for v := range c {
		fmt.Println("Received ", v)
		if v == 5 {
			fmt.Println("We got 5, and we're done.")

			// if the program didn't end here, the goroutine would still be running
			//time.Sleep(100 * time.Second)
			//break
		}
		fmt.Println(" v ", v)
	}
}
