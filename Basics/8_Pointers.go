package main

import (
	"fmt"
)

type gopher struct {
	age     int
	name    string
	isAdult bool
}

func main() {

	a := 12
	fmt.Println(*&a)
	fmt.Println(&a)
	fmt.Printf("type of a %T\n", a)
	fmt.Printf("type of *&a %T\n", *&a)

	b := &a
	fmt.Printf("type of b %T\n", b)

	fmt.Println(*b)
	*b = 23
	fmt.Println(a)

	ravi := gopher{name: "Ravi", age: 25}

	validateAge(&ravi)

	fmt.Println(ravi.isAdult)
}

func validateAge(g *gopher) {
	if g.age > 20 {
		g.isAdult = true
	} else {
		g.isAdult = false
	}
}
