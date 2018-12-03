package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context,c chan int) {

	i := 0
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			i++
			fmt.Println("generated : ", i)
			c <- i
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("ticket stopped")
		}
	}
}

func main() {

	ctx, _ := context.WithCancel(context.Background())

	c := make(chan int)

	go foo(c, ctx)

	for num := range c {
		fmt.Println("received : ", num)
		if num == 5 {
			//You can comment and see, the context is not being cancelled hence below for loop still receives the numbers.
			cancel()
			break
		}
	}

	for num := range c {
		fmt.Println("received 2 : ", num)
	}
}
