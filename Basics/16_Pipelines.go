package main

import (
	"fmt"
	"sync"
)



//Converting list of numbers into channels that emits the integers in the given list
//The gen function starts a go routine that sends the integers on the channel and closes the channel when all values have been sent
func gen(nums ...int) <-chan int {
	out := make(chan int, 2)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//sqr receives integers from a channel and returns a channel that emits square of each received integer.
func sqr(in <-chan int) <-chan int {
	out := make(chan int, 2)

	//Spanning a goroutine to square the number
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	// for _,c := range cs {
	// 	go func(c) {
	// 		for n := range c {
	// 			out <- n
	// 		}
	// 		wg.Done()
	// 	}()
	// }
	// wg.Add(len(cs))

	// go func() {
	// 	wg.Wait()
	// 	close(out)
	// }()

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	c := gen(2, 3, 4, 5, 6, 7, 8)
	out := sqr(c)

	fmt.Println("No Of Go routines	: ", runtime.NumGoroutine())
	fmt.Println("values returned : ", <-out)
	for n := range out {
		fmt.Println("values returned : ", n)
	}
	fmt.Println("No Of Go routines	: ", runtime.NumGoroutine())

}


//Fanout and Fanin
// func main() {
// 	in := gen(2, 3, 4, 5)
// 	//Distribute sqr work accross two goroutines

// 	//Fan out
// 	c1 := sqr(in)
// 	c2 := sqr(in)

// 	//Fan in
// 	//Consume the merged output from c1 and c2
// 	for n := range merge(c1, c2) {
// 		fmt.Println(n)
// 	}

// }

////Stopping short
// func main() {
// 	in := gen(2, 3, 4, 5)
// 	//Distribute sqr work accross two goroutines

// 	//Fan out
// 	c1 := sqr(in)
// 	c2 := sqr(in)

// 	//Fan in
// 	//Consume the merged output from c1 and c2
// 	out := merge(c1, c2)
// 	fmt.Println(<-out)

// }

// func main(){
// 	//Setup a done channel that is shared by whole pipeline,
// 	//and close that channel when the pipeline exits, as  a signal for all the goroutines we started to exit.

// 	done := make(chan struct{})

// 	defer close(done)

// 	in := gen(done,2,3,4)

// 	c1 := sqr(in)
// 	c2 := sqr(in)

// }
