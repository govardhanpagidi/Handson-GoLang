package main

import (
	"fmt"

	"github.com/dchest/uscrypt"
)

func main() {

	hash, _ := uscrypt.HashPassword(
		[]byte("Test@2019"),
		uscrypt.DefaultConfig,
	)
	fmt.Println("Hello, playground", hash)
}
