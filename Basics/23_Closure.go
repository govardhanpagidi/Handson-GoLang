package main

import "fmt"

func main() {
	NextInt := increment()
	fmt.Println(NextInt())
	fmt.Println(NextInt())
	fmt.Println(NextInt())
	fmt.Println(NextInt())

	NextInt2 := increment()
	fmt.Println(NextInt2())
	fmt.Println(NextInt2())
	fmt.Println(NextInt())

}

func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
