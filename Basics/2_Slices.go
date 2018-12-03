package main

import "fmt"

func main() {

	s := make([]string, 3)

	fmt.Println("emp:", s)

	s[0] = "1"
	s[1] = "2"
	s[2] = "3"

	s = append(s, "4")
	s = append(s, "5", "6")

	fmt.Println("emp list:", s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("copylist :", c)

	l := s[2:5]
	fmt.Println("sllice with 2:5", l)

	pr := s[:5]
	fmt.Println("preceeding with :5", pr)

	t := s[2:]
	fmt.Println("next with 2:", t)

}
