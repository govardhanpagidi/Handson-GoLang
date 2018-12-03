package main

import "fmt"
import "runtime"

func main() {
	fmt.Println("OS 			: ", runtime.GOOS)
	fmt.Println("Architecture 		: ", runtime.GOARCH)
	fmt.Println("No of CPUs		: ", runtime.NumCPU())
	fmt.Println("No Of Go routines	: ", runtime.NumGoroutine())
}
