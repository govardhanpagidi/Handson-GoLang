package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5))
	
	//passing slice as a variadic param
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Sum(slc...))
}

// Sum allows n nubmer of integers
func Sum(num ...int) int {
	sum := 0
	for _, num := range num {
		sum += num
	}
	return sum
}
