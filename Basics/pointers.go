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
