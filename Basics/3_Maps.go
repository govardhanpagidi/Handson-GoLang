package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["key1"] = 101
	m["key2"] = 102

	fmt.Println("map list", m)

	//Deleting an entry using key

	delete(m, "key1")

	fmt.Println("map after deletion", m)

	//Initialization while declaration

	m2 := map[string]int{"k1": 1001, "k2": 1002}
	fmt.Println(m2)

	m2["k3"] = 1003
	val,ok := m2["k4"]
	
	fmt.Println(val,ok)

}
