package main

import "fmt"

func main() {

	s := make([]string,4)

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("emp:", s)

	s[0] = "1"
	s[1] = "2"
	s[2] = "3"

	fmt.Println(s)
	test(s)
	fmt.Println(s)


fmt.Println(b)

	testArray(b)
	
	fmt.Println(b)



	fmt.Println(b)

	testArrayWithRef (&b)
	
	fmt.Println(b)
}


func test(s []string){
	s[0] = "234";
}

func testArray(a [5]int){
	a[0] = 67;
}


func testArrayWithRef(a *[5]int){
	a[0] = 78;
}
