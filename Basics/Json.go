package main
import "encoding/json"
import "fmt"

type response1 struct{
	Page int
	Fruits []string
}


type response2 struct{
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main(){

	 bolB,_ := json.Marshal(true)
	 fmt.Println(string(bolB))

	 intB,_ := json.Marshal(19)
	 fmt.Println(string(intB))

	 fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	
	slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	
	mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	
	
}