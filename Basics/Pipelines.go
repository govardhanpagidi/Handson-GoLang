package main

import "fmt"
import "runtime"

func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())


	c:= gen(2,3,4,5)
	out := sqr(c)
	for n := range out{
		fmt.Println("values returned : ",n)
	}

}

//Converting list of numbers into channels that emits the integers in the given list
//The gen function starts a go routine that sends the integers on the channel and closes the channel when all values have been sent
func gen(nums ...int ) <-chan int{
	out := make(chan int)
	
	go func(){
		for _,n := range nums{
			out <- n
		}
		close(out)
	}()
	return out
}

//sqr receives integers from a channel and returns a channel that emits square of each received integer.
func sqr(in <-chan int) <- chan int{
	out := make(chan int)

	go func(){
		for n := range in{
			out <- n*n
		}
		close(out)
	}()

	return out
}