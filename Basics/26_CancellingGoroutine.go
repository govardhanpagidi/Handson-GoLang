package main

import (
	"fmt"
	"time"
)

func f1(c chan int, quit chan bool) {
	i := 0
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			i++
			fmt.Println("Generated", i)
			c <- i
		case <-quit:
			ticker.Stop()
			fmt.Println("Ticker stopped. Closing channel.")
			close(c)
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
			q <- true
		}
		fmt.Println(" v ", v)
	}
}
