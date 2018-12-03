package main

import (
	"fmt"
)

func main() {

	//Arrays iteration using Range
	numbers := []int{1, 2, 3, 4, 5}

	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("sum is:", sum)

	//Maps iteration

	//Observation : For iteration, Not in order as we insert the elements
	for key, value := range map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5"} {

		fmt.Println("key , value ", key, value)
	}

	for key := range map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5"} {
		fmt.Println("keys ", key)
	}
}
