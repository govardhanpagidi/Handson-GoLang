package main
import "fmt"
import "strings"

type point struct {
    x, y int
}

func main(){
	p := point{1, 2}

	fmt.Printf("%v\n",p)

	fmt.Printf("%+v\n",p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf(" Displaying %s ", "stirng")

	fmt.Printf(" Display string with escapes %s ", "\"string\"")

	stngs := strings.Split("1,2,3,4,5,6,7,8,9",",")
	for _,s := range stngs {
		fmt.Println(s)
	}

	var r rune
	r=100
	fmt.Print("Printing rune ",r)

	//fmt.Print("Splitting a string : %v \n", stngs)
} 