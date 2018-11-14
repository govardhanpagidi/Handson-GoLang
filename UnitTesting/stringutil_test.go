package stringutil

import "testing"

//How to run
// Just cd to UnitTesting folder and run the below command, it identifies *_test.go files and run the unit tests
// > go test


func TestReverseString(t *testing.T){
	cases :=  []struct{
		in,want string
	}{
		{"Hello world","dlrow olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"testing","gniset"},
		{"", ""},
	}

	for _,c := range cases{
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse (%q) == %q, want %q", c.in,got,c.want)
		}
	}
}