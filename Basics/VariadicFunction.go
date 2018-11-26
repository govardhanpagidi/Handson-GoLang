package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func Sum(num ...int) int {
	sum := 0
	for _, num := range num {
		sum += num
	}
	return sum
}
