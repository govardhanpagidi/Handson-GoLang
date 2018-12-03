package main

import (
	"fmt"
	"errors"
)

func main(){

	val,err := sqrt(-1)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(val)
}

func sqrt (f float64) (float64, error){
	if f < 0 {
		return 0, errors.New("imporoper value to find Squre root")
	}
	return f,nil;
}