package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	//handling panic using deferred Recover.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func main() {
	firstName := "Govardhan"
	//lastName := "Pagidi"

	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
