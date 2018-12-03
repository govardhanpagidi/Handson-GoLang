package main

import "fmt"

func main() {
	var x int
	x++
	fmt.Println(x)

	i := c()
	fmt.Println(i)
}

func c() (j int) {
	defer func() { j++ }()
	return 3
}
