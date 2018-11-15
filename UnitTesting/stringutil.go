package stringutil

import (
	"fmt"
)

//This is just a string util method, to create an unit test case 

func Reverse(s string) (string){
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	if (s == ""){
		fmt.Print("nothing to reverse..")
	}
	return string(r)
}